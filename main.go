package main

import (
	"log"

	"fmt"

	"time"

	"crypto/sha1"

	"strconv"

	"strings"

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

//PostPersonAPI save data Person
func PostPersonAPI(ctx *iris.Context) {
	var person packages.Person
	person.Name = ctx.URLParam("name")
	person.Lastname = ctx.URLParam("last_name")
	person.CreateAt = time.Now()

	db := conection()

	db = db.Create(&person)

	fmt.Print(person.ID)

	if db.Error != nil {
		panic(db.Error)
	}
}

//PostUserAPI save data User
func PostUserAPI(ctx *iris.Context) {
	var user packages.User
	user.CI, _ = ctx.URLParamInt("ci")
	user.Password = GetSha1(ctx.URLParam("password"))
	user.CreateAt = time.Now()

	db := conection()

	db = db.Create(&user)

	fmt.Print(user.ID)

	if db.Error != nil {
		panic(db.Error)
	}
}

//GetSha1 password encryct
func GetSha1(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)

	aux := make([]string, len(bs))

	for i := range bs {
		aux[i] = strconv.Itoa(int(bs[i]))
	}

	return strings.Join(aux, "")
}

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

	iris.Post("/user/register", PostUserAPI)

	iris.Listen(":8080")
}
