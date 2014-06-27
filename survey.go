package main

import (
	"net/http"
 	"github.com/rebelnz/survey/tmpl"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.RenderHome(w, "index")
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	// http.Redirect(w, r, "/", 302)
	tmpl.RenderAccountHome(w)
}

func main() {
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/account",accountHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":9000",nil)
}











