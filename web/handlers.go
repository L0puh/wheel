package main

import (
	"net/http"
)

func home_handler(w http.ResponseWriter, r *http.Request){
	tmp.ExecuteTemplate(w, "home.html", nil)
}
