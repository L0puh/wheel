package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmp *template.Template

func main() {
	setup_static()	
	
	log.Print("Server is running on port 8080")
	http.HandleFunc("/", home_handler)
	http.ListenAndServe(":8080", nil)
}

func setup_static(){
	style := http.FileServer(http.Dir("static/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", style))
	var err error;
	tmp, err = template.ParseGlob("static/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
}
