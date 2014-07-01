package routes

import (
	c "github.com/rebelnz/survey/controllers"
	m "github.com/rebelnz/survey/models"
	"github.com/rebelnz/survey/db"
	"fmt"
	"html/template"
	"net/http"
)

func RenderHome(w http.ResponseWriter, tmpl string) {
	data := map[string]interface{}{"Tmpl":tmpl}
	t, _ := template.ParseFiles("templates/_fore.html", "templates/home.html")
	if err := t.Execute(w, data); err != nil {
		fmt.Println(err)
		return
	}
}

func RenderRegister(w http.ResponseWriter, r *http.Request, tmpl string) {
	data := m.Page{Tmpl: tmpl}
	if r.Method == "POST" {
		err := c.Register(r) // checks username isnt taken then registers the user
		if err != nil {
			fmt.Println(err)
			data.Flash = "Username Taken"
		} else {
			http.Redirect(w, r, "/account", 302)
		}
	}
	t, _ := template.ParseFiles("templates/_fore.html", "templates/register.html")
	if err := t.Execute(w, data); err != nil {
		fmt.Println(err)
		return
	}
}

func RenderAccount(w http.ResponseWriter, r *http.Request, tmpl string) {
	person := m.Person{}	
	p := db.DB.First(&person).Value
	data := map[string]interface{}{"Tmpl":tmpl,"Person":p}
	t, _ := template.ParseFiles("templates/_aft.html", "templates/account.html")
	if err := t.Execute(w, data); err != nil {
		fmt.Println(err)
		return
	}
}

