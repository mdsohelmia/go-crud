package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sohelcse1999/gocrud/controllers"
	dbclient "github.com/sohelcse1999/gocrud/db_client"
)

func main() {

	dbclient.InitilizeDbConnection()
	r:=gin.Default();

	r.POST("/posts",controllers.CreatePost)
	r.GET("/posts",controllers.GetPosts)

	if err := r.Run(":5000"); err !=nil {
		panic(err.Error())
	}

}