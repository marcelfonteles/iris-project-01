package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-project-01/app/models"
	"strconv"
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
	b.Handle("POST", "/new", "CreateClient")
	b.Handle("GET", "/{id:uint}/delete", "DeleteClient")
	b.Handle("GET", "/{id:uint}/edit", "EditClient")
}

func (c *clientController) Index() mvc.Result {
	var clients []models.Client
	clients = models.GetClients()
	return mvc.View{
		Name: "clients/index",
		Data: clients,
	}
}

func (c *clientController) Form(ctx iris.Context) mvc.Result {
	var users []models.User
	users = models.GetUsers()
	ctx.ViewData("users", users)
	//ctx.ViewData()
	return mvc.View{
		Name: "clients/form",
	}
}

func (c *clientController) CreateClient(ctx iris.Context) {
	var client models.Client
	userID, _ := strconv.Atoi(ctx.FormValue("user_id"))
	client = models.Client{
		Model:     gorm.Model{},
		FirstName:  ctx.FormValue("first_name"),
		LastName:  ctx.FormValue("last_name"),
		Area:      ctx.FormValue("area"),
		User_id:   uint(userID),
	}
	models.NewClient(&client)
	ctx.Redirect("/clients")
}

func (c *clientController) DeleteClient(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	models.DeleteClient(id)
	ctx.Redirect("/clients")
}

func (c *clientController) EditClient(ctx iris.Context) mvc.Result {
	id, _ := ctx.Params().GetUint("id")
	client := models.GetClient(id)
	users := models.GetUsers()
	ctx.ViewData("FirstName", client.FirstName)
	ctx.ViewData("LastName", client.LastName)
	ctx.ViewData("Area", client.Area)
	ctx.ViewData("client", client)
	ctx.ViewData("users", users)
	return mvc.View{
		Name: "clients/form",
	}

}



