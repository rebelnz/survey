package controllers

import (
	"github.com/rebelnz/survey/db"
	m "github.com/rebelnz/survey/models"
	"code.google.com/p/go.crypto/bcrypt"
	"net/http"
	"log"
	"os"
	"fmt"
	"errors"
)

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

func Register(r *http.Request) (error) {
	person := m.Person{}
	username, password := r.FormValue("username"), r.FormValue("password")
	p := db.DB.Where("username = ?",username).Find(&person)
	if p.RecordNotFound() {
		// username doesnt exist? - register new user
		pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		db.DB.Save(m.Person{Username:username,Password:string(pwd)})
		fmt.Println("new person saved")
		return nil	
	} else {
		err := errors.New("username exists")
		return err
	}
	return nil	
}



















