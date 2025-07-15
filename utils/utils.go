package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gomarkdown/markdown"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// MarkdownToHTML converts markdown string to HTML
func MarkdownToHTML(md string) string {
	html := markdown.ToHTML([]byte(md), nil, nil)
	return string(html)
}
