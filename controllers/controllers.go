package controllers

import (
	"github.com/rebelnz/survey/db"
	m "github.com/rebelnz/survey/models"
	// s "github.com/rebelnz/survey/session"
	"code.google.com/p/go.crypto/bcrypt"
	"net/http"
	"log"
	"os"
	"fmt"
	"errors"
	"time"
)

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

func Register(w http.ResponseWriter, r *http.Request) (error) {
	account := m.Account{}
	username, email, password := r.FormValue("username"), r.FormValue("email"), r.FormValue("password")
	// check for duplicate
	p := db.DB.Where("username = ? or email = ?",username, email).Find(&account)
	
	if p.RecordNotFound() {
		// username doesnt exist? - register new user
		pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		db.DB.Save(m.Account{
			Username: username,
			Email: email,
			Password: string(pwd),
			CreatedAt: time.Now(),
		})
		fmt.Println("new account saved")
		return nil	
	} else {
		err := errors.New("Email and/or Username already exists")
		return err
	}
	return nil	
}
