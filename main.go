package main

import (
	"log"

	"./packages"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
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
	db, e := gorm.Open("postgres", "host=localhost user=postgres dbname=ReportCollision sslmode=disable password=k3yl0gg3r")
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
	return db
}

//Migrations for create tables
func Migrations() {
	db := conection()

	//db.Rollback()
	db.AutoMigrate(&packages.User{}, &packages.Person{}, &packages.Climate{}, &packages.Severity{}, &packages.Factors{}, &packages.Accident{}, &packages.Photo{}, &packages.Involved{}, &packages.AccidentFactors{})
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
