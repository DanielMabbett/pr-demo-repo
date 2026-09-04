package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jumpstart-demo/pr-demo-repo/internal/status"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8091"
	}

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, status.Report("pr-demo-api", true))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status":"ok"}`)
	})

	log.Printf("pr-demo-repo listening on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
