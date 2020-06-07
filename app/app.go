package app

import (
	"log"
	connection "yalla_go/db"
	"yalla_go/product"
	"yalla_go/user"

	"github.com/gin-gonic/gin"
)

// Run to run app in main
func Run() {
	db := connection.Connect()
	err := db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	defer connection.Close(db)
	app := gin.Default()
	app.Group("/api")
	UserAPI := InitUserApi(db)
	ProductAPI := InitProductApi(db)
	user.Routes(app, UserAPI)
	product.Routes(app, ProductAPI)
	app.Run()
}
