package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello from Go + ArgoCD 👋")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	type HelloResponse struct {
		Msg string `json:"msg"`
	}
	json.NewEncoder(w).Encode(HelloResponse{"Hello, Matt!"})
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", handleIndex)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
