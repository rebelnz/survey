package main

import (
	"net/http"
 	"github.com/rebelnz/survey/routes"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	routes.RenderHome(w,"home")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	routes.RenderRegister(w, r, "register")
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	routes.RenderAccount(w, r, "account")
}

func main() {
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/account",accountHandler)
	http.HandleFunc("/register",registerHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":9000",nil)
}











