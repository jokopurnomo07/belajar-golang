package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	
	r.GET("/", Home)
	r.GET("/get", getting)
	r.GET("/get/:name", getYourName)
	r.POST("/add-data", addData)
	r.PUT("/edit-data/:id", editData)
	r.POST("/upload", uploadFile)
	r.Run()
}