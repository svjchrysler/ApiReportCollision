package main

import (
	"log"

	"fmt"

	"./packages"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
)

const (
	//HOST database
	HOST = "localhost"
	//USER database
	USER = "postgres"
	//DB database
	DB = "ReportCollision"
	//SSL database
	SSL = "disable"
	//PASSWORD database
	PASSWORD = "k3yl0gg3r"
)

/*
func GetApiIris(ctx *iris.Context) {

	db := conection()

	var rest Person

	db.Table("people").Find(&rest)

	err := ctx.JSON(iris.StatusOK, rest)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}*/

//GetAPI data for Android
func GetAPI(ctx *iris.Context) {
	user := packages.User{CI: 9684203, Password: "hola"}
	err := ctx.JSON(iris.StatusOK, user)
	if err != nil {
		log.Fatal(err)
	}
}

/*func PostPerson(ctx *iris.Context) {
	var person Person
	name := ctx.URLParam("name")
	db := conection()
	person.Name = name
	db = db.Create(&person)
	if db.Error != nil {
		panic(db.Error)
	}
}*/

func conection() *gorm.DB {
	conexion := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", HOST, USER, DB, SSL, PASSWORD)
	db, e := gorm.Open("postgres", conexion)
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
	return db
}

//Migrations for create tables
func Migrations() {
	db := conection()

	if !db.HasTable(&packages.Weather{}) {
		db.CreateTable(&packages.Weather{})
	}

	if !db.HasTable(&packages.Involved{}) {
		db.CreateTable(&packages.Involved{})
	}

	if !db.HasTable(&packages.Severity{}) {
		db.CreateTable(&packages.Severity{})
	}

	if !db.HasTable(&packages.User{}) {
		db.CreateTable(&packages.User{})
	}

	if !db.HasTable(&packages.Person{}) {
		db.CreateTable(&packages.Person{}).
			AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	}

	if !db.HasTable(&packages.Accident{}) {
		db.CreateTable(&packages.Accident{}).
			AddForeignKey("weather_id", "weathers(id)", "CASCADE", "CASCADE").
			AddForeignKey("person_id", "people(id)", "CASCADE", "CASCADE").
			AddForeignKey("severity_id", "severities(id)", "CASCADE", "CASCADE")
	}

	if !db.HasTable(&packages.Factors{}) {
		db.CreateTable(&packages.Factors{})
	}

	if !db.HasTable(&packages.AccidentFactors{}) {
		db.CreateTable(&packages.AccidentFactors{}).
			AddForeignKey("accident_id", "accidents(id)", "CASCADE", "CASCADE").
			AddForeignKey("factors_id", "factors(id)", "CASCADE", "CASCADE")
	}

	if !db.HasTable(&packages.Photo{}) {
		db.CreateTable(&packages.Photo{}).
			AddForeignKey("accident_id", "accidents(id)", "CASCADE", "CASCADE")
	}

	if !db.HasTable(&packages.AccidentInvolved{}) {
		db.CreateTable(&packages.AccidentInvolved{}).
			AddForeignKey("accident_id", "accidents(id)", "CASCADE", "CASCADE").
			AddForeignKey("involved_id", "involveds(id)", "CASCADE", "CASCADE")
	}

	defer db.Close()

}

func main() {

	Migrations()

	/*iris.Get("/", GetApiIris)*/

	iris.Get("/api", GetAPI)

	/*	iris.Post("/person", PostPerson)*/

	/*iris.Get("/", func(c *iris.Context) {
		c.JSON(iris.StatusOK, iris.Map{
			"Name":     "Iris",
			"Released": "13 March 2016",
		})
	})*/

	iris.Listen(":8080")
}
