package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name" form:"name"`
	Age int `json:"age" form:"age"`
}

func getting(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "getting your data....",
    })
}

func Home(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "This home baby",
	})
}

func getYourName(c *gin.Context){
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"your_name": name,
	})
}

func addData(c *gin.Context){
	var form Person
	err := c.Bind(&form)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Zaki error",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"name": form.Name,
			"age": form.Age,
		})
	}
}

func editData(c *gin.Context){
	id := c.Param("id")
	var form Person
	err := c.ShouldBind(&form)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "u have an error",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"id": id,
			"nama": form.Name,
			"age": form.Age,
		})
	}
}

func uploadFile(c *gin.Context){
	file, _ := c.FormFile("upload")
	var form Person
	err := c.ShouldBind(&form)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "u have an error",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"your_file": file.Filename,
			"name": form.Name,
			"age": form.Age,
		})
	}
}