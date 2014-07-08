package main

import (
	"net/http"
	// "fmt"
	"github.com/gorilla/mux"
	"github.com/rebelnz/survey/routes"
	// m "github.com/rebelnz/survey/middleware" // DEV Use function - not being used 
	c "github.com/rebelnz/survey/controllers"
)

var router = mux.NewRouter()

func indexHandler(w http.ResponseWriter, r *http.Request) {
	routes.RenderHome(w, r, "home")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	routes.RenderRegister(w, r, "register")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	routes.RenderLogin(w, r, "login")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	c.Logout(w,r)
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	routes.RenderAccount(w, r, "account")
}
	
func main() {
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/register", registerHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/logout", logoutHandler)
	// router.HandleFunc("/account", m.Use(accountHandler, c.RequireLogin))
	router.HandleFunc("/account", accountHandler)
	http.Handle("/", router)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":9000", nil)
}
