package main

import "github.com/gin-gonic/gin"

type Person struct{
	Name string;
	Age int
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		var p Person

		p.Name = "Sunil Kumar";
		p.Age = 24;

		c.Bind(&p)
         
		c.JSON(200, gin.H{
			"name": p.Name,
			"Age":p.Age,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}