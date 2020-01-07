package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sinmetal/gaeimage/backend"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	http.HandleFunc("/v2/", backend.ImageHandlerV2)
	http.HandleFunc("/v1/", backend.ImageHandler)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), http.DefaultServeMux); err != nil {
		log.Printf("failed ListenAndServe err=%+v", err)
	}
}
