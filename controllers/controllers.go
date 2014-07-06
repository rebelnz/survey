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
)

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

func Register(w http.ResponseWriter, r *http.Request) (error) {
	account := m.Account{}
	username, password := r.FormValue("username"), r.FormValue("password")
	p := db.DB.Where("username = ?",username).Find(&account)
	if p.RecordNotFound() {
		// username doesnt exist? - register new user
		pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		db.DB.Save(m.Account{Username:username,Password:string(pwd)})
		fmt.Println("new account saved")
		return nil	
	} else {
		err := errors.New("username exists")
		return err
	}
	return nil	
}
