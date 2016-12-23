package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
)

type User struct {
	Name     string
	Lastname string
}

type Person struct {
	Id   int64 `gorm:"primary_key"`
	Name string
}

func GetApiIris(ctx *iris.Context) {
	/*slice := make(map[nil]User)
	slice = append(slice, User{Name: "Jhon", Lastname: "Salguero"})
	slice = User{Name: "luis", Lastname: "Salguero"}
	slice = User{Name: "xavier", Lastname: "Salguero"}*/

	db := conection()

	var rest Person

	db.Table("people").Find(&rest)

	/*user := User{Name: "luis", Lastname: "Salguero"}*/
	/*userT := User{Name: "Juan", Lastname: "Salguero"}*/

	err := ctx.JSON(iris.StatusOK, rest)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func GetApi(ctx *iris.Context) {
	person := Person{Name: "Angel"}
	err := ctx.JSON(iris.StatusOK, person)
	if err != nil {
		log.Fatal(err)
	}
}

func PostPerson(ctx *iris.Context) {
	var person Person
	name := ctx.URLParam("name")
	db := conection()
	person.Name = name
	db = db.Create(&person)
	if db.Error != nil {
		panic(db.Error)
	}
}

func conection() *gorm.DB {
	db, e := gorm.Open("postgres", "host=localhost user=postgres dbname=ReportCollision sslmode=disable password=k3yl0gg3r")
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
	return db
}

func Migrations() {
	db := conection()
	db.AutoMigrate(&User{}, &Person{})
	defer db.Close()
}

func main() {

	Migrations()

	iris.Get("/", GetApiIris)

	iris.Get("/api", GetApi)

	iris.Post("/person", PostPerson)

	/*iris.Get("/", func(c *iris.Context) {
		c.JSON(iris.StatusOK, iris.Map{
			"Name":     "Iris",
			"Released": "13 March 2016",
		})
	})*/

	iris.Listen(":8080")
}
