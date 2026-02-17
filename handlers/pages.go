package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Parse template once
var tmpl = template.Must(template.ParseFiles("templates/page.html"))

// PageData holds info for template rendering
type PageData struct {
	PageNumber int
	Color      string
	HasNext    bool
	HasPrev    bool
	NextPage   int
	PrevPage   int
}

// Map page number to background color
var colors = map[int]string{
	1: "#FF4C4C", // Red
	2: "#4C6FFF", // Blue
	3: "#4CFF88", // Green
	4: "#B84CFF", // Purple
}

// PageHandler serves pages 1-4
func PageHandler(w http.ResponseWriter, r *http.Request) {
	page := 1

	pageParam := chi.URLParam(r, "number")
	if pageParam != "" {
		p, err := strconv.Atoi(pageParam)
		if err == nil && p >= 1 && p <= 4 {
			page = p
		}
	}

	data := PageData{
		PageNumber: page,
		Color:      colors[page],
		HasNext:    page < 4,
		HasPrev:    page > 1,
		NextPage:   page + 1,
		PrevPage:   page - 1,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}
