package main

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie/models"
	"groupie/services"
)

func main() {
	http.HandleFunc("/dates", models.DatesHandler)
	http.HandleFunc("/locations", models.LocationsHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/artist/", models.ArtistDetailsHandler)

	http.HandleFunc("/search", models.SearchHandler)
	http.HandleFunc("/searchLocation", models.SearchLocationHandler)
	fmt.Println("Server is starting...")
	fmt.Println("Go on http://localhost:8080/")
	fmt.Println("To shut down the server press CTRL + C")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, artists)
}
