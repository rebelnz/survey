package routes

import (
 	"github.com/rebelnz/survey/models"
	"github.com/rebelnz/survey/db"
 	"html/template"
	"net/http"
	"fmt"
	"time"
	// "reflect"
)

var person = models.Person{}

func RenderHome(w http.ResponseWriter, tmpl string) {	
	data := models.Page{Tmpl:tmpl}
	t, _ := template.ParseFiles("templates/_fore.html", "templates/home.html")
	if err := t.Execute(w,data); err != nil {
		fmt.Println(err)
		return
	}
}

func RenderRegister(w http.ResponseWriter, tmpl string) {		
	data := models.Page{Tmpl:tmpl}
	t, _ := template.ParseFiles("templates/_fore.html", "templates/register.html")
	if err := t.Execute(w,data); err != nil {
		fmt.Println(err)
		return
	}
}

func RenderAccount(w http.ResponseWriter) {
	xdata := db.DB.First(&person).Value
	fmt.Println(xdata)
	accData := models.Account{AccountName:"Rebel Account"}
	data := struct {
		AccountData models.Account
		Ts time.Time
	} {
		accData,
		time.Now(),
	}
	t, _ := template.ParseFiles("templates/_aft.html","templates/account.html")		
	if err := t.Execute(w, data); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
