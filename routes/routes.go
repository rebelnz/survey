package routes

import (
	con "github.com/rebelnz/survey/controllers"
	mod "github.com/rebelnz/survey/models"
	"github.com/rebelnz/survey/db"
	"html/template"
	"net/http"
	"fmt"
)

func RenderHome(w http.ResponseWriter, r *http.Request, title string) {
	data := mod.Page{
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
	data := mod.Page{
		Title: title,
		CSS: []string{"register","style"},
	}
	if r.Method == "POST" {
		err := con.Register(w,r) // checks username isnt taken then registers the user
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

func RenderLogin(w http.ResponseWriter, r *http.Request, title string) {
	t, _ := template.ParseFiles("templates/_front.html", "templates/login.html")
	data := mod.Page{
		Title: title,
		CSS: []string{"login"}}
	if r.Method == "POST" {
		err := con.Login(w, r)
		if err != nil {
			errorString := []string{err.Error(),"There was a problem logging in"}
			data.Message = append(data.Message, errorString...)
			t.Execute(w, data)
		} else {
			fmt.Println("success") // DEV
			http.Redirect(w, r, "/account", 302)
		}
	}
	if err := t.Execute(w, data); err != nil {
		fmt.Println(err)
		return
	}
}


func RenderAccount(w http.ResponseWriter, r *http.Request, title string) {
	session, _ := con.Store.Get(r, con.SESSION_NAME)
	userid,_ := session.Values["UserID"]
	// fmt.Println("User ID", userid) // DEV
	if userid == nil {
		http.Redirect(w, r, "/login", 302)		
	}
	account := db.DB.First(&mod.Account{},userid).Value // DEV this only gets first account
	// fmt.Println(account) // DEV
	data := mod.Page{
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

