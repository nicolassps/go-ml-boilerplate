package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/linear"
)

const modelPath = "./models/model.json"

func main() {
	if len(os.Args) < 3 || os.Args[1] != "--mode" {
		fmt.Println("Uso: go run . --mode [train|serve]")
		return
	}

	mode := os.Args[2]

	switch mode {
	case "train":
		train()
	case "serve":
		serve()
	default:
		fmt.Println("Modo inválido. Use 'train' ou 'serve'.")
	}
}

func train() {
	fmt.Println("🔄 Treinando o modelo...")

	file, err := os.Open("./data/champions.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var X [][]float64
	var y []float64

	for i, record := range records {
		if i == 0 {
			continue
		}

		homeShots, _ := strconv.ParseFloat(record[11], 64)
		homeSOT, _ := strconv.ParseFloat(record[13], 64)
		homeFouls, _ := strconv.ParseFloat(record[15], 64)
		homeCorners, _ := strconv.ParseFloat(record[17], 64)
		fthGoals, _ := strconv.ParseFloat(record[4], 64)

		X = append(X, []float64{homeShots, homeSOT, homeFouls, homeCorners})
		y = append(y, fthGoals)
	}

	X = normalizeFeatures(X)

	model := linear.NewLeastSquares(base.BatchGA, 1e-6, 0.0, 1000, X, y)

	if err := model.Learn(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Treinamento concluído!")

	if err := saveModel(model, modelPath); err != nil {
		log.Fatal(err)
	}

	fmt.Println("📂 Modelo salvo em", modelPath)
}

func serve() {
	fmt.Println("🚀 Servindo a API...")

	params, err := loadModel(modelPath)
	if err != nil {
		log.Fatal("Erro ao carregar o modelo:", err)
	}

	model := linear.NewLeastSquares(base.BatchGA, 1e-6, 0.0, 1000, nil, nil)
	model.Parameters = params

	http.HandleFunc("/predict", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		var input struct {
			HomeShots   float64 `json:"home_shots"`
			HomeSOT     float64 `json:"home_sot"`
			HomeFouls   float64 `json:"home_fouls"`
			HomeCorners float64 `json:"home_corners"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Erro ao processar JSON", http.StatusBadRequest)
			return
		}

		features := []float64{input.HomeShots, input.HomeSOT, input.HomeFouls, input.HomeCorners}
		prediction, err := model.Predict(features)
		if err != nil {
			http.Error(w, "Erro ao fazer previsão", http.StatusInternalServerError)
			return
		}

		response := map[string]float64{"predicted_goals": prediction[0]}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	log.Println("📡 API rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func normalizeFeatures(X [][]float64) [][]float64 {
	normalized := make([][]float64, len(X))
	for i := range X {
		normalized[i] = make([]float64, len(X[i]))
	}

	for j := 0; j < len(X[0]); j++ {
		var sum float64
		for i := 0; i < len(X); i++ {
			sum += X[i][j]
		}
		mean := sum / float64(len(X))

		var variance float64
		for i := 0; i < len(X); i++ {
			variance += (X[i][j] - mean) * (X[i][j] - mean)
		}
		stdDev := math.Sqrt(variance / float64(len(X)))

		for i := 0; i < len(X); i++ {
			normalized[i][j] = (X[i][j] - mean) / stdDev
		}
	}

	return normalized
}

func saveModel(model *linear.LeastSquares, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(model.Parameters)
}

func loadModel(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var params []float64
	if err := json.NewDecoder(file).Decode(&params); err != nil {
		return nil, err
	}

	return params, nil
}
