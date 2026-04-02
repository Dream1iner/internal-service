package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func envHandler(w http.ResponseWriter, r *http.Request) {
	vars := make(map[string]string)
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		vars[parts[0]] = parts[1]
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(vars)
}

func main() {
	http.HandleFunc("/", envHandler)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
