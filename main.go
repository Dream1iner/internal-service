package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

const pageTemplate = `<!DOCTYPE html>
<html>
<head>
  <title>Environment Variables</title>
  <style>
    body { font-family: monospace; background: #1e1e2e; color: #cdd6f4; margin: 2rem; }
    h1 { color: #89b4fa; }
    table { border-collapse: collapse; width: 100%%; max-width: 900px; }
    th, td { text-align: left; padding: 0.5rem 1rem; border-bottom: 1px solid #313244; }
    th { color: #a6adc8; }
    td:first-child { color: #f38ba8; }
    td:last-child { color: #a6e3a1; word-break: break-all; }
    .badge { display: inline-block; padding: 0.2rem 0.6rem; border-radius: 4px; font-size: 0.85rem; }
    .prod { background: #f38ba8; color: #1e1e2e; }
    .dev { background: #a6e3a1; color: #1e1e2e; }
  </style>
</head>
<body>
  <h1>internal-service %s</h1>
  <table>
    <tr><th>Variable</th><th>Value</th></tr>
    %s
  </table>
</body>
</html>`

func getEnvVars() map[string]string {
	vars := make(map[string]string)
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		vars[parts[0]] = parts[1]
	}
	return vars
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(getEnvVars())
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := getEnvVars()

	badge := `<span class="badge dev">dev</span>`
	if vars["PROD"] == "true" {
		badge = `<span class="badge prod">prod</span>`
	}

	keys := make([]string, 0, len(vars))
	for k := range vars {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var rows strings.Builder
	for _, k := range keys {
		rows.WriteString("<tr><td>" + k + "</td><td>" + vars[k] + "</td></tr>\n    ")
	}

	w.Header().Set("Content-Type", "text/html")
	page := strings.Replace(pageTemplate, "%s", badge, 1)
	page = strings.Replace(page, "%s", rows.String(), 1)
	w.Write([]byte(page))
}

func main() {
	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/json", jsonHandler)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
