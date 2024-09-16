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
	CUI        string         `json:"cui"`
	Nume       string         `json:"nume"`
	NrRegCom   string         `json:"nr_reg_com"`
	Adresa     string         `json:"adresa"`
	Value      string         `json:"value"`
	Cod        string         `json:"cod"`
	Judet      string         `json:"judet"`
	Comuna     sql.NullString `json:"comuna"`
	Localitate sql.NullString `json:"localitate"`
}

type FirmResponse struct {
	CUI        string `json:"cui"`
	Nume       string `json:"nume"`
	NrRegCom   string `json:"nr_reg_com"`
	Adresa     string `json:"adresa"`
	Value      string `json:"value"`
	Cod        string `json:"cod"`
	Judet      string `json:"judet"`
	Comuna     string `json:"comuna"`
	Localitate string `json:"localitate"`
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
	row := db.QueryRow("SELECT * FROM date_firme_nume_cui WHERE cui = ?", cui)

	var firm Firm
	err := row.Scan(&firm.CUI, &firm.Nume, &firm.NrRegCom, &firm.Adresa, &firm.Value, &firm.Cod, &firm.Judet, &firm.Comuna, &firm.Localitate)
	if err == sql.ErrNoRows {
		return FirmResponse{}, nil
	} else if err != nil {
		return FirmResponse{}, err
	}

	resp := FirmResponse{
		CUI:        firm.CUI,
		Nume:       firm.Nume,
		NrRegCom:   firm.NrRegCom,
		Adresa:     firm.Adresa,
		Value:      firm.Value,
		Cod:        firm.Cod,
		Judet:      firm.Judet,
		Comuna:     firm.Comuna.String,
		Localitate: firm.Localitate.String,
	}

	return resp, nil
}
