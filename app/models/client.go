package models

import (
	"github.com/jinzhu/gorm"
	"iris-project-01/datasource"
)

type Client struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(30)"`
	LastName  string `gorm:"type:varchar(30)"`
	Area      string `gorm:"type:varchar(30)"`
	User_id   uint   `gorm:"type:integer"`
}

func GetClients(client *[]Client) *[]Client {
	db := datasource.Connection()
	db.AutoMigrate(&Client{})
	db.Find(client)
	return client
}

func GetClient(client *Client, id uint) *Client {
	db := datasource.Connection()
	db.AutoMigrate(&Client{})
	db.Find(client, id)
	return client
}


