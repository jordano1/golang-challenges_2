package main

import (
	"html/template"
	"log"
	"net/http"
)

// #1 uncomment line 10
// type pizza int
// #2 using pizza, we would use replace "jordan" on line 14 with:
// (d pizza) ServeHTTP

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

func main() {
	// can create a type of int (alternate approach) ( see #1 )
	// #3 then declare d of type pizza here, and pass it into handle
	// var d pizza
	// and pass it into handle, otherwise you could just use the handlerfunc method
	http.Handle("/jordan/", http.HandlerFunc(jordan))
	http.ListenAndServe(":8080", nil)
}
