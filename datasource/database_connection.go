package datasource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	)

func Connection() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=marcelvieira dbname=go_database password=j55fonteles")
	if err != nil {
		fmt.Println("Could not connect to database")
	}
	db.LogMode(true)
	return db
}
