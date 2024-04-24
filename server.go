package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	ID  string `json:"id"`
	Ml  string `json:"ml,omitempty"`
}

var db = make(map[string]string)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.Ml != "" {
		db[data.ID] = data.Ml
		fmt.Fprintf(w, "ID %s hat den Füllstand %s ml\n", data.ID, data.Ml)
	} else {
		if ml, ok := db[data.ID]; ok {
			fmt.Fprintf(w, "ID %s hat den Füllstand %s ml\n", data.ID, ml)
		} else {
			http.Error(w, "ID nicht gefunden", http.StatusNotFound)
		}
	}
}

func addData(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db[data.ID] = data.Ml
	fmt.Fprintf(w, "ID %s mit Füllstand %s ml wurde hinzugefügt\n", data.ID, data.Ml)
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/add", addData)
	http.ListenAndServe(":8080", nil)
}