package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/lib/pq"
)

var DB gorm.DB

func init() {
	var err error

	DB, err = gorm.Open("postgres", "user=rebel password=Ohlolk70 dbname=survey port=5433 sslmode=disable")

	if err != nil {
		panic(fmt.Sprintf("Error: ",err))
	}
}



















