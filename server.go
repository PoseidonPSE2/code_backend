package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	ID        string `json:"id"`
	Ml        string `json:"ml,omitempty"`
	WaterType string `json:"waterType,omitempty"`
}

type DatabaseEntry struct {
	Ml        string
	WaterType string
}

var db = make(map[string]DatabaseEntry)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("Fehler beim Dekodieren der Anfrage: %v", err)
		http.Error(w, "Ungültige Anfrage", http.StatusBadRequest)
		return
	}

	log.Printf("Anfrage erhalten für ID: %s", data.ID)

	if entry, ok := db[data.ID]; ok {
		response := fmt.Sprintf("ID %s hat den Füllstand %s ml und Wasserart %s\n", data.ID, entry.Ml, entry.WaterType)
		log.Printf("Antwort gesendet: %s", response)
		fmt.Fprint(w, response)
	} else {
		log.Printf("ID nicht gefunden: %s", data.ID)
		http.Error(w, "ID nicht gefunden", http.StatusNotFound)
	}
}

func addData(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("Fehler beim Dekodieren der Anfrage: %v", err)
		http.Error(w, "Ungültige Anfrage", http.StatusBadRequest)
		return
	}

	db[data.ID] = DatabaseEntry{Ml: data.Ml, WaterType: data.WaterType}
	response := fmt.Sprintf("ID %s mit Füllstand %s ml und Wasserart %s wurde hinzugefügt\n", data.ID, data.Ml, data.WaterType)
	log.Printf("Daten hinzugefügt: %s", response)
	fmt.Fprint(w, response)
}

func addDataManually(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ml := r.URL.Query().Get("ml")
	waterType := r.URL.Query().Get("waterType")

	if id == "" || ml == "" || waterType == "" {
		http.Error(w, "Bitte geben Sie id, ml und Wasserart an", http.StatusBadRequest)
		return
	}

	db[id] = DatabaseEntry{Ml: ml, WaterType: waterType}
	response := fmt.Sprintf("ID %s mit Füllstand %s ml und Wasserart %s wurde hinzugefügt\n", id, ml, waterType)
	log.Printf("Daten hinzugefügt: %s", response)
	fmt.Fprint(w, response)
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/add", addData)
	http.HandleFunc("/addManually", addDataManually)
	log.Println("Server läuft und hört auf Port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server konnte nicht gestartet werden: %v", err)
	}
}
