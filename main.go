package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello from Go + ArgoCD 👋")
}

func main() {
	http.HandleFunc("/", handleIndex)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
