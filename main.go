package main

// Iris Framework with Gorm (ORM)

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-project-01/app/controllers"
)



func main() {
	app := iris.New()
	mvc.Configure(app.Party("/"), controllers.User_controller)
	app.Logger().SetLevel("debug")

	app.RegisterView(iris.HTML("app/views", ".html"))

	app.Run(iris.Addr(":3000"))
}
