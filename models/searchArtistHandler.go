package models

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"groupie/services"
)

// Add function that fetch data,process search query and renders result to index.html
func SearchHandler(w http.ResponseWriter, r *http.Request) {


	query := r.URL.Query().Get("query")// accepting user input



	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filteredArtists []services.Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	tmpl, err := template.ParseFiles(filepath.Join("templates", "index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, filteredArtists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
