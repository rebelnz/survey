package routes

import (
	c "github.com/rebelnz/survey/controllers"
	m "github.com/rebelnz/survey/models"
	// s "github.com/rebelnz/survey/session"
	"github.com/rebelnz/survey/db"
	"html/template"
	"net/http"
	"fmt"
)

func RenderHome(w http.ResponseWriter, r *http.Request, title string) {
	data := m.Page{
		Title: title,
		UserID:1,
	}
	t, _ := template.ParseFiles("templates/_front.html", "templates/home.html")
	if err := t.Execute(w, data); err != nil {
		fmt.Println(err)
		return
	}
}

func RenderRegister(w http.ResponseWriter, r *http.Request, title string) {
	data := m.Page{
		Title: title,
		CSS: []string{"register","style"},
	}
	if r.Method == "POST" {
		err := c.Register(w,r) // checks username isnt taken then registers the user
		if err != nil {
			errorString := []string{err.Error(),"Error 2"}
			data.Message = append(data.Message, errorString...)
		} else {
			http.Redirect(w, r, "/account", 302)
		}
	}
	t, _ := template.ParseFiles("templates/_front.html", "templates/register.html")
	if err := t.Execute(w, data); err != nil {
		fmt.Println(err)
		return
	}
}

func RenderAccount(w http.ResponseWriter, r *http.Request, title string) {
	account := db.DB.First(&m.Account{}).Value
	data := m.Page{
		Title: title,
		Account: account,
		CSS: []string{"account"},
	}
	t, _ := template.ParseFiles("templates/_back.html", "templates/account.html")
	if err := t.Execute(w, data); err != nil {
		fmt.Println(err)
		return
	}
}

