package main

// Iris Framework with Gorm (ORM)

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris/hero"
	"log"
	"strconv"
)

type User struct {
	gorm.Model
	first_name string //`gorm:"type:varchar(30);column:first_name"`
	last_name  string //`gorm:"type:varchar(30);column:last_name"`
	sex        string //`gorm:"type:enum('M', 'F', 'O');column:sex"`
	age        int    //`gorm:"type:smallint;column:age"`
}

func newApp() (*iris.Application, *gorm.DB) {
	app := iris.New()
	// Database Connection
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=marcelvieira dbname=go_database password=j55fonteles")
	if err != nil {
		log.Fatal("Could not logon on database.", err)
	} else {
		log.Println("Logon on database successfully")
	}
	// Log Mode
	db.LogMode(true)
	// Closing connection when program ends.
	// defer db.Close()
	// Creating tables when their do not exists
	fmt.Println("Running migrations....")
	db.AutoMigrate(&User{})
	fmt.Println("Running migrations....[OK]")

	// Register Views
	app.RegisterView(iris.HTML("app/views", ".html"))

	return app, db
}

func main() {
	app, db := newApp()
	defer db.Close()
	formHandler := hero.Handler(form)
	app.Get("/", formHandler)

	app.Post("/form/post", func(ctx iris.Context) {
		fname := ctx.FormValue("first_name")
		lname := ctx.FormValue("last_name")
		age, _ := strconv.Atoi(ctx.FormValue("age"))
		sex := ctx.FormValue("sex")
		user := User{
			first_name: fname,
			last_name:  lname,
			age:        age,
			sex:        sex,
		}
		fmt.Println(fname, lname, age, sex)
		fmt.Println(user)
		if err := db.Create(&user); err != nil {
			log.Println("Could not create a new user.")
			ctx.Redirect("/")
			return
		}
		ctx.JSON(iris.Map{
			"code": 200,
		})

	})

	app.Run(iris.Addr(":3000"))
}

func form(ctx iris.Context) {
	ctx.View("index.html")
}
