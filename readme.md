# Go Machine Learning Boilerplate

## English Version 🇬🇧

### Overview
This repository is a **boilerplate** for beginners in Go who want to explore **Machine Learning**. It provides a simple implementation of a **linear regression model** using the `goml` package to predict football match outcomes based on historical data.

### How to Run

1. **Clone the repository**
   ```sh
   git clone https://github.com/nicolassps/go-ml-boilerplate.git
   cd go-ml-boilerplate
   ```
2. **Ensure Go is installed** (version 1.18 or later recommended).

3. **Prepare your dataset**
   - Place a CSV file in `./data/champions.csv` with the following structure:

     ```csv
     Date,Season,HomeTeam,AwayTeam,FTH Goals,FTA Goals,FT Result,HTH Goals,HTA Goals,HT Result,Referee,H Shots,A Shots,H SOT,A SOT,H Fouls,A Fouls,H Corners,A Corners,H Yellow,A Yellow,H Red,A Red,Display_Order,League
     16/01/2025,2024/25,Ipswich Town,Brighton & Hove Albion,0,2,A,0,1,A,T Harrington,5,11,3,5,13,14,1,9,2,2,0,0,20250116,Premier League
     ```

4. **Train the model**
   ```sh
   go run . --mode train
   ```
   This will read the CSV file, train the model, and save the parameters in `model.json`.

5. **Start the API server**
   ```sh
   go run . --mode serve
   ```
   The server will run on `http://localhost:8080`.

### Making Predictions

Send a **POST request** to `http://localhost:8080/predict` with JSON input:

```json
{
  "home_shots": 10,
  "home_sot": 5,
  "home_fouls": 8,
  "home_corners": 4
}
```

Example using `curl`:
```sh
curl -X POST http://localhost:8080/predict \
     -H "Content-Type: application/json" \
     -d '{"home_shots":10, "home_sot":5, "home_fouls":8, "home_corners":4}'
```

Expected response:
```json
{
  "predicted_goals": 2.34
}
```

---

## Versão em Português 🇧🇷

### Visão Geral
Este repositório é um **boilerplate** para iniciantes em Go que desejam explorar **Machine Learning**. Ele fornece uma implementação simples de um **modelo de regressão linear** usando o pacote `goml` para prever resultados de jogos de futebol com base em dados históricos.

### Como Rodar

1. **Clone o repositório**
   ```sh
   git clone https://github.com/nicolassps/go-ml-boilerplate.git
   cd go-ml-boilerplate
   ```
2. **Certifique-se de ter o Go instalado** (recomendado: versão 1.18 ou superior).

3. **Prepare seu conjunto de dados**
   - Coloque um arquivo CSV em `./data/champions.csv` com a seguinte estrutura:

     ```csv
     Date,Season,HomeTeam,AwayTeam,FTH Goals,FTA Goals,FT Result,HTH Goals,HTA Goals,HT Result,Referee,H Shots,A Shots,H SOT,A SOT,H Fouls,A Fouls,H Corners,A Corners,H Yellow,A Yellow,H Red,A Red,Display_Order,League
     16/01/2025,2024/25,Ipswich Town,Brighton & Hove Albion,0,2,A,0,1,A,T Harrington,5,11,3,5,13,14,1,9,2,2,0,0,20250116,Premier League
     ```

4. **Treine o modelo**
   ```sh
   go run . --mode train
   ```
   Isso lerá o arquivo CSV, treinará o modelo e salvará os parâmetros em `model.json`.

5. **Inicie o servidor da API**
   ```sh
   go run . --mode serve
   ```
   O servidor rodará em `http://localhost:8080`.

### Fazendo Previsões

Envie uma **requisição POST** para `http://localhost:8080/predict` com um JSON como este:

- home_shots: O número de chutes realizados pela equipe da casa (home).
- home_sot: O número de chutes ao gol feitos pela equipe da casa (home). "SOT" é uma sigla para "Shots On Target".
- home_fouls: O número de faltas cometidas pela equipe da casa.
- home_corners: O número de escanteios (corners) conquistados pela equipe da casa.

```json
{
  "home_shots": 10,
  "home_sot": 5,
  "home_fouls": 8,
  "home_corners": 4
}
```

Exemplo usando `curl`:
```sh
curl -X POST http://localhost:8080/predict \
     -H "Content-Type: application/json" \
     -d '{"home_shots":10, "home_sot":5, "home_fouls":8, "home_corners":4}'
```

Resposta esperada:
```json
{
  "predicted_goals": 2.34
}
```

---

🚀 **Agora você está pronto para experimentar Machine Learning com Go!**

