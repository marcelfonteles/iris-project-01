package controllers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-project-01/app/models"
	"strconv"
)

func User_controller(app *mvc.Application) {
	app.Handle(new(userController))
}

type userController struct {
	User models.User
	Users []models.User
}

func (u *userController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/users/form", "Form")
	b.Handle("GET", "/users/{id:uint}", "GetUser")
	b.Handle("GET", "/users", "GetUsers")
	b.Handle("POST", "/users/new", "NewUser")
}

func (u *userController) Form() mvc.Result {
	return mvc.View{
		Name: "form/form",
	}
}

func (u *userController) GetUser(id uint) mvc.Result {
	var user models.User
	user = models.GetUser(id)
	fmt.Println(user)
	return mvc.View{
		Name: "users/show",
		Data: user,
	}
}

func (u *userController) GetUsers() mvc.Result {
	var users []models.User
	users = models.GetUsers()
	fmt.Println(users)
	return mvc.View{
		Name: "users/show",
		Data: users,
	}
}

func (u *userController) NewUser(ctx iris.Context) {
	var user models.User
	age, _ := strconv.Atoi(ctx.FormValue("age"))
	user = models.User{
		Model:     gorm.Model{},
		FirstName: ctx.FormValue("first_name"),
		LastName:  ctx.FormValue("last_name"),
		Sex:       ctx.FormValue("sex"),
		Age:       age,
	}
	id := models.NewUser(user)
	if id != 0 {
		ctx.Redirect("/users/" + strconv.Itoa(int(id)))
	} else {
		ctx.Redirect("/users/form")
	}
	fmt.Println(user)
}

