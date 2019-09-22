package models

import (
	"github.com/jinzhu/gorm"
	"iris-project-01/datasource"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(30);column:first_name"`
	LastName  string `gorm:"type:varchar(30);column:last_name"`
	Sex       string `gorm:"type:char(1);column:sex"`
	Age       int    `gorm:"type:smallint;column:age"`
}

func GetUser(id uint) []User {
	var user []User
	db := datasource.Connection()
	defer db.Close()
	db.First(&user, id)
	return user
}

func GetUsers() []User {
	var users []User
	db := datasource.Connection()
	defer db.Close()
	defer db.AutoMigrate(&User{})
	db.Order("id").Find(&users)
	return users
}

func NewUser(user User) uint {
	db := datasource.Connection()
	db.Create(&user)
	return user.ID
}

func DeleteUser(id uint) bool {
	db := datasource.Connection()
	var user User
	db.Where("id = ?", id).Find(&user)
	if user.ID != 0 {
		db.Delete(&user)
		return true
	} else {
		return false
	}
}
