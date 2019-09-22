package controllers

import (
	"github.com/kataras/iris/mvc"
	"iris-project-01/app/models"
)

func Client_controller(app *mvc.Application) {
	app.Handle(new(clientController))
}

type clientController struct {
	Client models.Client
	Clients []models.Client
}

func (c *clientController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "Index")
	b.Handle("GET", "/new", "Form")
}

func (c *clientController) Index() mvc.Result {
	var clients []models.Client
	models.GetClients(&clients)
	return mvc.View{
		Name: "clients/index",
		Data: clients,
	}
}

func (c *clientController) Form() mvc.Result {
	var users []models.User
	users = models.GetUsers()
	return mvc.View{
		Name: "clients/form",
		Data: users,
	}
}



