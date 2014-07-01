package controllers

import (
	"github.com/rebelnz/survey/db"
	m "github.com/rebelnz/survey/models"
	"code.google.com/p/go.crypto/bcrypt"
	"net/http"
	"log"
	"os"
	"fmt"
)

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)


func Register(r *http.Request) (error) {
	person := m.Person{}
	username, password := r.FormValue("username"), r.FormValue("password")
	p := db.DB.Where("username = ?",username).Find(&person)
	if p.RecordNotFound() {
		pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			return err
		}
		db.DB.Save(m.Person{Username:username,Password:string(pwd)})
		fmt.Println("new person saved")
		return nil	
	} else {
		fmt.Println("username exists")
	}
	return nil	
}



















