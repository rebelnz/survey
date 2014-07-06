package db

import (
	"github.com/jinzhu/gorm"
	_"github.com/lib/pq"
	"fmt"
)

var DB gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", "user=rebel password=Ohlolk70 dbname=survey port=5433 sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("Database connection error: ", err))
	}
}



















