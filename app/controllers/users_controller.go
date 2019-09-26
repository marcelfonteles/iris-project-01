package controllers

import (
	"fmt"
	_ "fmt"
	"github.com/golang/protobuf/proto"
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
	b.Handle("GET", "/form", "Form")
	b.Handle("GET", "/{id:uint}", "GetUser")
	b.Handle("GET", "/", "GetUsers")
	b.Handle("POST", "/new", "NewUser")
	b.Handle("GET", "/{id:uint}/delete", "DeleteUser")
	b.Handle("GET", "/{id:uint}/edit", "EditUser")

	// TODO: change for a new controller later
	b.Handle("GET", "/persons", "Persons")
	b.Handle("POST", "/persons/new", "NewPersons")
}

func (u *userController) Form() mvc.Result {
	return mvc.View{
		Name: "users/form",
	}
}

func (u *userController) GetUser(id uint) mvc.Result {
	var user models.User
	user = models.GetUser(id)
	return mvc.View{
		Name: "users/show",
		Data: user,
	}
}

func (u *userController) GetUsers() mvc.Result {
	var users []models.User
	users = models.GetUsers()
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
}

func (u *userController) DeleteUser(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	success := models.DeleteUser(id)
	if success {
		ctx.Redirect("/users")
	} else {
		ctx.Redirect("/users/" + strconv.Itoa(int(id)))
	}
}

func (u * userController) EditUser(ctx iris.Context) mvc.Result {
	id, _ := ctx.Params().GetUint("id")
	user := models.GetUser(id)
	return mvc.View{
		Name: "users/form",
		Data: user,
	}
}

// TODO: change for a new conroller later
func (u *userController) Persons(ctx iris.Context) {
	var people []models.Person
	people = models.GetPersons()
	//ctx.JSON(people)
	var data []byte
	//for _, value := range people {
	//	proto_data, _ := proto.Marshal(&value)
	//	data = append(data, proto_data)
	//}
	data, _ = proto.Marshal(&people[0])
	fmt.Println(data)
	fmt.Printf("%T\n",data)
	ctx.Binary(data)
}

func (u *userController) NewPersons(ctx iris.Context) {
	data, _ := ctx.GetBody()
	fmt.Println(data)
	fmt.Println(string(data))
	var person models.Person
	proto.Unmarshal(data, &person)
	fmt.Println(person)
}
