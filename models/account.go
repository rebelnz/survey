package models

import (
	"time"
	"github.com/rebelnz/survey/db"
	"errors"
	"database/sql"
)

var ErrNoUsername = errors.New("Username doesnt exist")

type Account struct {
	Id        int64
	Username  string `sql:"size:255;not null"`
	Password  string `sql:"size:255;not null"`
	Firstname string `sql:"size:255"`
	Lastname  string `sql:"size:255"`
	Email     string `sql:"size:255;not null"`
	Priv      int64  `sql:"not null; default=1"`
	CreatedAt time.Time //TODO default time now 
	UpdatedAt time.Time //TODO default time now 
	// LastAccessIP string
	// LastAccessTime time.Time
}


func GetAccountByUsername(username string) (Account, error) {
	a := Account{}
	err := db.DB.Where("username = ?", username).First(&a).Error
	if err != sql.ErrNoRows {
		return a, ErrNoUsername
	} else if err != nil {
		return a, err
	}
	return a, nil
}

// Authenticate checks if a cookie is a valid login
// func (a *Account) Authenticate(r *http.Request) bool {
// 	session, _ := store.Get(r, "auth-cookie")
// 	id, ok := session.Values["Id"]
// 	if ! ok {
// 		fmt.Println("account.authenticate = false")
// 		return false
// 	}
// 	var err error
// 	a.Id, err = strconv.ParseInt(id.(string), 10, 0)
// 	fmt.Print(p.Id)
// 	if err != nil {
// 		return false
// 	}
// 	fmt.Println("account.authenticate = true")
// 	return true
// }









