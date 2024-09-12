package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type ApiConfig struct {
	Port int `json:"port"`
}

type Firm struct {
	CUI           string         `json:"cui"`
	Nume          string         `json:"nume"`
	NrRegCom      sql.NullString `json:"nr_reg_com"`
	NumereTelefon sql.NullString `json:"numere_telefon"`
	Adresa        string         `json:"adresa"`
	CodJudetAprox string         `json:"cod_judet_aprox"`
}

type FirmResponse struct {
	CUI           string `json:"cui"`
	Nume          string `json:"nume"`
	NrRegCom      string `json:"nr_reg_com"`
	NumereTelefon string `json:"numere_telefon"`
	Adresa        string `json:"adresa"`
	CodJudetAprox string `json:"cod_judet_aprox"`
}

func main() {

	serverPort := getApiConfig()

	db, err := sql.Open("sqlite3", "./date_firme_cui.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/date-cui/", func(w http.ResponseWriter, r *http.Request) {
		cui := r.URL.Path[len("/date-cui/"):]
		firm, err := getFirmByCUI(db, cui)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(firm)
	})

	fmt.Println("Server is running on port ", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}

func getApiConfig() string {

	file, err := os.Open("./datecuiapi.config.json")
	if err != nil {
		return "3240"
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var config ApiConfig
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	portStr := fmt.Sprintf("%d", config.Port)
	fmt.Println("Port as string:", portStr)

	return portStr

}

func getFirmByCUI(db *sql.DB, cui string) (FirmResponse, error) {
	row := db.QueryRow("SELECT cui, nume, nr_reg_com, numere_telefon, adresa, cod_judet_aprox FROM date_firme_cui WHERE cui = ?", cui)

	var firm Firm
	err := row.Scan(&firm.CUI, &firm.Nume, &firm.NrRegCom, &firm.NumereTelefon, &firm.Adresa, &firm.CodJudetAprox)
	if err == sql.ErrNoRows {
		return FirmResponse{}, nil
	} else if err != nil {
		return FirmResponse{}, err
	}

	resp := FirmResponse{
		CUI:           firm.CUI,
		Nume:          firm.Nume,
		NrRegCom:      firm.NrRegCom.String,
		NumereTelefon: firm.NumereTelefon.String,
		Adresa:        firm.Adresa,
		CodJudetAprox: firm.CodJudetAprox,
	}

	return resp, nil
}
