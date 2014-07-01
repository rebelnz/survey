package models

import (
	"github.com/rebelnz/survey/db"
	"errors"
	"log"
	"time"
	"os"
)

var err error
var ErrUsernameTaken = errors.New("username already taken")
var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

// database

type Account struct {
	Id          int64
	PersonId    int64
	AccountName string `sql:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Campaign struct {
	Id        int64
	AccountId int64
	Account   string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Survey struct {
	Id         int64
	AccountId  int64
	SurveyName string `sql:"size:255"`
	Questions  string `sql:"type:json"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Question struct {
	// incomplete
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Answer struct {
	Id         int64
	SurveyId   int64
	QuestionId int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Survey_question struct {
	// many to many
	SurveyId   int64
	QuestionId int64	
}

// meta
type Page struct {
	Tmpl string
	Flash string
}


func init() {
	db.DB.SingularTable(true)
	db.DB.AutoMigrate(Person{})
	db.DB.AutoMigrate(Account{})
	db.DB.AutoMigrate(Campaign{})
	db.DB.AutoMigrate(Survey{})
	db.DB.AutoMigrate(Question{})
	db.DB.AutoMigrate(Answer{})
	db.DB.AutoMigrate(Survey_question{})
	log.Println("Tables created")
	
	// p := Person{
	// 	Username: "rebel",
	// 	Password: "pass",
	// }
	// err = db.DB.Save(&p).Error
	// if err != nil {
	// 	Logger.Println(err)
	// }
}




















