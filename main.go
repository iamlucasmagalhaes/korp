package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Project struct {
	Name string `json:"nome"`
	Time string `json:"horario"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UTC().Format(time.RFC3339)
	response := Project{
		Name: "Projeto Korp",
		Time: now,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/projeto-korp", handler)
	http.ListenAndServe(":8080", nil)
}
