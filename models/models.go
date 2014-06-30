package models

import (
	"errors"
	"github.com/rebelnz/survey/db"
	"log"
	"time"
	"os"
)

var err error
var ErrUsernameTaken = errors.New("username already taken")
var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

type Person struct {
	Id        int64
	Username  string `sql:"size:255;not null"`
	Password  string `sql:"size:255;not null"`
	Firstname string `sql:"size:255"`
	Lastname  string `sql:"size:255"`
	Email     string `sql:"size:255;not null"`
	Priv      int64  `sql:"not null; default=1"`
	CreatedAt time.Time 
	UpdatedAt time.Time
}

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
	
	init_user := Person{
		Username: "admin",
		Password: "password",
	}

	err = db.DB.Save(&init_user).Error
	if err != nil {
		Logger.Println(err)
	}

}

type Page struct {
	Tmpl string
}
