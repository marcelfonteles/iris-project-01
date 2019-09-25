package models

import (
	"iris-project-01/datasource"
)

func GetPersons() []Person {
	db := datasource.Connection()
	defer db.Close()
	var people []Person
	db.AutoMigrate(&Person{})
	db.Find(&people)
	return people
}