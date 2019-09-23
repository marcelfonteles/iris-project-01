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
	UserFirstName string `gorm:"-"`
}


func GetClients() []Client {
	var clients []Client
	db := datasource.Connection()
	defer db.Close()
	db.AutoMigrate(&Client{})
	db.Order("id").Find(&clients)
	for key, value := range clients {
		var user User
		db.First(&user, value.User_id)
		clients[key].UserFirstName = user.FirstName
	}
	return clients
}

func GetClient(id uint) Client {
	var client Client
	var user User
	db := datasource.Connection()
	defer db.Close()
	db.AutoMigrate(&Client{})
	db.Find(&client, id)
	db.Find(&user, client.User_id)
	client.UserFirstName = user.FirstName
	return client
}

func NewClient(client *Client) {
	db := datasource.Connection()
	defer db.Close()
	db.Create(client)
}

func DeleteClient(clientID uint) *Client {
	db := datasource.Connection()
	defer db.Close()
	var client Client
	db.First(&client, clientID)
	db.Delete(&client)

	return &client
}


