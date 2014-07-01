package models

import (
	"time"
)


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

