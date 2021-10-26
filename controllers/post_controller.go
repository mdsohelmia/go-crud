package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dbclient "github.com/sohelcse1999/gocrud/db_client"
)

type Post struct {
	Id			int64 `json:"id"`
	Name        string `json:"name"`
	UserId      int64  `json:"user_id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Slug        string `json:"slug"`
}

func CreatePost(c *gin.Context) {
	var requestBody Post

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "Invalid respone Request boyd.",
			"description":err.Error(),
		})
		return
	}
	res,erro :=dbclient.DbClient.Exec("INSERT INTO posts(name,user_id,description,status,slug) VALUES(?,?,?,?,?)",
		requestBody.Name,
		requestBody.UserId,
		requestBody.Description,
		requestBody.Status,
		requestBody.Slug,
	)
	if erro !=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}
	id,err :=res.LastInsertId()
		if err !=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"id":id,
	})
}

func GetPosts(c *gin.Context){

	var posts[] Post;

	rows,err := dbclient.DbClient.Query("SELECT id,name,user_id,description,status,slug FROM posts ORDER BY id desc LIMIT 1000");

	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}
	for rows.Next(){
		var singlePost Post;
		if err:=rows.Scan(&singlePost.Id,&singlePost.Name,&singlePost.UserId,&singlePost.Description,&singlePost.Status,&singlePost.Slug);err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message":err.Error(),
		})
		return
		}

		posts = append(posts, singlePost)
	}
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"data":posts,
	})
}