package routes

import (
 	m "github.com/rebelnz/survey/models"
 	c "github.com/rebelnz/survey/controllers"
	// "github.com/rebelnz/survey/db"
 	"html/template"
	"net/http"
	"fmt"
	// "time"
	// "reflect"
)

var person = m.Person{}

func RenderHome(w http.ResponseWriter, tmpl string) {	
	data := m.Page{Tmpl:tmpl}
	t, _ := template.ParseFiles("templates/_fore.html", "templates/home.html")
	if err := t.Execute(w,data); err != nil {
		fmt.Println(err)
		return
	}
}

func RenderRegister(w http.ResponseWriter,r *http.Request, tmpl string) {		
	if r.Method == "POST" {
		
		err := c.Register(r) // check username isnt taken then register the user
		
		if err != nil {
			fmt.Println("ok")
		}

		// r.ParseForm()
		// fmt.Println(r.FormValue("username"))
		// for k, v := range r.Form {
		// 	fmt.Println("key: ", k)
		// 	fmt.Println("value: ", v)
		// }
		http.Redirect(w, r, "/account", 302)
	}
	
	data := m.Page{Tmpl:tmpl}
	t, _ := template.ParseFiles("templates/_fore.html", "templates/register.html")
	if err := t.Execute(w,data); err != nil {
		fmt.Println(err)
		return
	}
}

func RenderAccount(w http.ResponseWriter) {
	// xdata := db.DB.First(&person).Value
	// fmt.Println(xdata)
	// accData := m.Account{AccountName:"Rebel Account"}
	// data := struct {
	// 	AccountData m.Account
	// 	Ts time.Time
	// } {
	// 	accData,
	// 	time.Now(),
	// }
	t, _ := template.ParseFiles("templates/_aft.html","templates/account.html")		
	// if err := t.Execute(w, data); err != nil {
	if err := t.Execute(w, nil); err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(data)
}
