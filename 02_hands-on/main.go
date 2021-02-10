package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/jordan/", jordan)
	http.ListenAndServe(":8080", nil)
}

func jordan(w http.ResponseWriter, req *http.Request) {

	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad"}

	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("err", err)
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", sages)
	if err != nil {
		log.Fatalln("err", err)
	}
}
