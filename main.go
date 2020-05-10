package main

import (
	"fmt"
	"log"
	connection "yalla_go/db"
	"yalla_go/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db := connection.Connect()
	err := db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	defer connection.Close(db)
	router := gin.Default()
	UserAPI := InitUserApi(db)
	user.Routes(router, UserAPI)
	router.Run()
	fmt.Println("connection established")
}
