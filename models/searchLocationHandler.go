package models

import (
	"html/template"
	"net/http"
	"path/filepath"
	"slices"
	"strings"

	"groupie/services"
)

// Add function that fetch data,process search query and renders result to index.html
func SearchLocationHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query") // accepting user input

	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filteredArtists []services.Artist
	var artistNames []string
	for _, artist := range artists {
		for _, location := range artist.Locations {
			if strings.Contains(strings.ToLower(location), strings.ToLower(query)) {

				if slices.Contains(artistNames, artist.Name) {
					continue
				}
				filteredArtists = append(filteredArtists, artist)
				artistNames = append(artistNames, artist.Name)
			}
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
