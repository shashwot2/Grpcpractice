package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.gohtml")
	})
	fmt.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(fmt.Println(err))
	}
}

func render(w http.ResponseWriter, t string) {
	partials := []string{
		"./cmd/web/templates/index.gohtml",
		"./cmd/web/templates/footer.gohtml",
		"./cmd/web/templates/header.gohtml",
	}
	var templateSlice []string

	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t))
	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
