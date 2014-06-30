package tmpl

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

func RenderHome(w http.ResponseWriter) {	
	t, _ := template.ParseFiles("templates/index.html")
	data := db.DB.First(&person).Value
	if err := t.Execute(w,data); err != nil {
		fmt.Println(err)
		return
	}
}
func RenderRegister(w http.ResponseWriter) {	
	t, _ := template.ParseFiles("templates/register.html")		
	if err := t.Execute(w,nil); err != nil {
		fmt.Println(err)
		return
	}
}

func RenderAccount(w http.ResponseWriter) {
	accData := models.Account{AccountName:"Rebel Account"}
	data := struct {
		AccountData models.Account
		Ts time.Time
	} {
		accData,
		time.Now(),
	}
	t, _ := template.ParseFiles("templates/_base.html","templates/account.html")		
	if err := t.Execute(w, data); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
