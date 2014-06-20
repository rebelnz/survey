package tpl

import (
 	"html/template"
	"net/http"
)


type Page struct {
    Title string
    Body  []byte
}

func RenderHome(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles("templates/" + tmpl + ".html")
    t.Execute(w, nil)
}


// func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
//     err := templates.ExecuteTemplate(w, tmpl + ".html", p)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
// }

	
// var SurveysTpl = template.Must(template.ParseFiles(
// 	"templates/_base.html",
// 	"templates/surveys/index.html",
// ))

// var AccountTpl = template.Must(template.ParseFiles(
// 	"templates/_base.html",
// 	"templates/account/index.html",
// ))










