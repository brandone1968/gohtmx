package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println("Ciao Go")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseGlob("templates/*.html"))

		data := struct {
			Name        string
			Title       string
			Description string
			Socials     map[string]string
			Features    []string
		}{
			Name:        "Gopher",
			Title:       "Welcome to Go",
			Description: "This is a simple web application using Go and HTML templates.",
			Socials: map[string]string{
				"Facebook":  "https://www.facebook.com",
				"Twitter/X": "https://www.twitter.com",
				"LinkedIn":  "https://www.linkedin.com",
				"Instagram": "https://www.instagram.com",
				"GitHub":    "https://www.github.com",
			},
			Features: []string{
				"Fast and efficient",
				"Easy to learn",
				"Strongly typed",
			},
		}
		err := tmpl.ExecuteTemplate(w, "home.html", data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":3000", nil)
}
