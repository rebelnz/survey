package main

import (
	"net/http"
 	"github.com/rebelnz/survey/tpls"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.RenderHome(w, "index")
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	// http.Redirect(w, r, "/", 302)
	tpl.RenderAccountHome(w)
}

// func surveysHandler(w http.ResponseWriter, r *http.Request) {
// 	s1 := Survey{SurveyName:"Mar_102",SurveyProject:"Marketing"} 
// 	if err := SurveysTpl.Execute(w, s1); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func accountHandler(w http.ResponseWriter, r *http.Request) {
// 	u := User{UserName:"Rebel"}
// 	if err := AccountTpl.Execute(w, u); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

func main() {
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/account",accountHandler)
	// http.HandleFunc("/account",accountHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":9000",nil)
}

