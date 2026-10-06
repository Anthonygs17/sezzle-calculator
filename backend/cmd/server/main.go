package main

import (
	"log"
	"net/http"

	"github.com/Anthonygs17/sezzle-calculator/backend/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/calculate", handler.Calculate)

	log.Println("server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
