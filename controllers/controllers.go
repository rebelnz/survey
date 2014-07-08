package controllers

import (
	"github.com/rebelnz/survey/db"
	mod "github.com/rebelnz/survey/models"
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/gorilla/sessions"
	"net/http"
	"log"
	"os"
	"errors"
	"time"
	"fmt"
)

const (
	ENCRYPTSECRET = "Ohk8ahjaeya5sCHANGEMEoocatah8UoB"
	SESSION_NAME = "survey-session"
)

var Store = sessions.NewCookieStore([]byte(ENCRYPTSECRET))

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)


func Register(w http.ResponseWriter, r *http.Request) (error) {
	account := mod.Account{}
	username, email, password := r.FormValue("username"), r.FormValue("email"), r.FormValue("password")
	// check for duplicate
	p := db.DB.Where("username = ? or email = ?",username, email).Find(&account)
	if p.RecordNotFound() {
		// username doesnt exist? - register new user
		pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		db.DB.Save(mod.Account{
			Username: username,
			Email: email,
			Password: string(pwd),
			CreatedAt: time.Now(),
		})
		return nil	
	} else {
		err := errors.New("Email and/or Username already exists")
		return err
	}
	return nil	
}


func Login(w http.ResponseWriter, r *http.Request) (error) {
	username, password := r.FormValue("username"), r.FormValue("password")
	a, err := mod.GetAccountByUsername(username)
	if err != nil && err != mod.ErrNoUsername {
		return  err
	}
	err = bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	if err != nil {
		return errors.New("Incorrect password")
	} else {
		session, err := Store.Get(r, SESSION_NAME)
		if err != nil {
			fmt.Println(err)
		}
		session.Options = &sessions.Options{
			Path: "/",
		}
		session.Values["UserID"] = a.Id
		session.Values["Auth"] = 1
		session.Save(r,w)
		return nil
	}
	return errors.New("Could not login")
 }

func Logout(w http.ResponseWriter, r *http.Request) (error) {
	session, _ := Store.Get(r, SESSION_NAME) // DEV handle error
	delete(session.Values,"UserID")
	delete(session.Values,"Auth")
	http.Redirect(w, r, "/login", 302)
	return nil
}


func RequireLogin(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, SESSION_NAME)
		fmt.Println(session)
		// 	http.Redirect(w, r, "/login", 302)
		// }
		handler.ServeHTTP(w, r)
	}
}	









