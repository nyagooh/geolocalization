package models

import (
	"html/template"
	"net/http"
	"path/filepath"
	"slices"
	"strings"

	"groupie/services"
)

// Add function that filters data based on date fetch data,process search query and renders result to index.html
func SearchDatesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query") // accepting user input

	query = formatDateInput(query)

	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filteredArtists []services.Artist
	var artistNames []string
	for _, artist := range artists {
		for _, dates := range artist.Relations {
			for _, date := range dates {
				// fmt.Println(date)
				if strings.Contains(strings.ToLower(date), strings.ToLower(query)) {

					if slices.Contains(artistNames, artist.Name) {
						continue
					}
					filteredArtists = append(filteredArtists, artist)
					artistNames = append(artistNames, artist.Name)
				}
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

func formatDateInput(s string) string {
	splitDate := strings.Split(s, "-")
	return splitDate[2] + "-" + splitDate[1] + "-" + splitDate[0]
}
