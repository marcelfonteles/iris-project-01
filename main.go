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
	FirstName string `gorm:"type:varchar(30);column:first_name"`
	LastName  string `gorm:"type:varchar(30);column:last_name"`
	Sex       string `gorm:"type:char(1);column:sex"`
	Age       int    `gorm:"type:smallint;column:age"`
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

	// Saving user to database
	app.Post("/form/post", func(ctx iris.Context) {
		fname := ctx.FormValue("first_name")
		lname := ctx.FormValue("last_name")
		age, _ := strconv.Atoi(ctx.FormValue("age"))
		sex := ctx.FormValue("sex")
		user := User{
			FirstName: fname,
			LastName:  lname,
			Age:       age,
			Sex:       sex,
		}
		fmt.Println(fname, lname, age, sex)
		if err := db.Create(&user); err != nil {
			log.Println("Could not create a new user.")
			ctx.Redirect("/")
			return
		}
		ctx.JSON(iris.Map{
			"code": 200,
		})

	})
	// Listing all users from database
	app.Get("/users/{id:uint}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetUint("id")
		var user User
		db.First(&user, int(id))
		fmt.Println(user)
		id     = user.ID
		fname := user.FirstName
		lname := user.LastName
		age   := user.Age
		sex   := user.Sex
		fmt.Println(id, fname, lname, age, sex)
		if id != 0 {
			ctx.ViewData("id", strconv.Itoa(int(id)))
			ctx.ViewData("fname", fname)
			ctx.ViewData("lname", lname)
			ctx.ViewData("age", strconv.Itoa(age))
			ctx.ViewData("sex", sex)
		} else {
			ctx.ViewData("fname", "Could not find a user.")
		}
		ctx.View("user.html")
	})

	app.Run(iris.Addr(":3000"))
}

func form(ctx iris.Context) {
	ctx.View("index.html")
}
