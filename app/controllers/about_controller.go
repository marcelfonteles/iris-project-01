package controllers

import (
	"github.com/kataras/iris/mvc"
)

func About_controller (app *mvc.Application) {
	app.Handle(new(aboutController))
}

type aboutController struct {

}

func (a *aboutController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "About")
}

func (a *aboutController) About() mvc.Result {
	return mvc.View{
		Name: "about/about",
	}
}
