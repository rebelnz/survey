package tmpl

import (
 	"github.com/rebelnz/survey/models"
 	"html/template"
	"net/http"
	"fmt"
	"time"
)

var person = models.Person{}

func RenderHome(w http.ResponseWriter, tmpl string) {	

	data := models.Page{Title:"Survey Home Title"}
	t, _ := template.ParseFiles("templates/" + tmpl + ".html")

	// data := db.DB.First(person)
	
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
