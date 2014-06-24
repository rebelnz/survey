package tpl

import (
 	"github.com/rebelnz/survey/models"
 	"html/template"
	"net/http"
	"fmt"
	"time"
)

// var AccountTpl = template.Must(template.ParseFiles(
// 	"templates/_base.html",
// 	"templates/account/index.html",
// ))

// var SurveysTpl = template.Must(template.ParseFiles(
// 	"templates/_base.html",
// 	"templates/surveys/index.html",
// ))

func RenderHome(w http.ResponseWriter, tmpl string) {	

	data := models.Page{Title:"Survey Home Title"}
	t, _ := template.ParseFiles("templates/" + tmpl + ".html")	

	if err := t.Execute(w,data); err != nil {
		fmt.Println(err)
		return
	}

}

func RenderAccountHome(w http.ResponseWriter) {
	accData := models.Account{AccountName:"Rebel Account"}
	data := struct {
		AccountData models.Account
		Ts time.Time
	} {
		accData,
		time.Now(),
	}
	t, _ := template.ParseFiles("templates/_base.html","templates/account/index.html")		
	if err := t.Execute(w, data); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}


	











