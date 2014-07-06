package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rebelnz/survey/routes"
	m "github.com/rebelnz/survey/middleware"
)

var router = mux.NewRouter()

func indexHandler(w http.ResponseWriter, r *http.Request) {
	routes.RenderHome(w, r, "home")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	routes.RenderRegister(w, r, "register")
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	routes.RenderAccount(w, r, "account")
}

func requireLogin(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("requires login")
		handler.ServeHTTP(w, r)
	}
}	
	
func main() {
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/account", m.Use(accountHandler, requireLogin))
	router.HandleFunc("/register", registerHandler)
	http.Handle("/", router)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":9000", nil)
}

















