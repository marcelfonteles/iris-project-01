package main

// Iris Framework with Gorm (ORM)

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-project-01/app/controllers"
)



func main() {
	app := iris.New()
	// Configuring Controllers
	mvc.Configure(app.Party("/users"), controllers.User_controller)
	mvc.Configure(app.Party("/"), controllers.About_controller)
	mvc.Configure(app.Party("/clients"), controllers.Client_controller)
	// Log appear on cosole
	app.Logger().SetLevel("debug")
	// Configuring Views
	app.RegisterView(iris.HTML("app/views", ".html"))
	// Running Application
	app.Run(iris.Addr(":3000"))
}
