package main

import "github.com/gin-gonic/gin"

type Person struct{
	Name string;
	Age int
}

func main() {
	r := gin.Default() // Creates a gin router with default middleware:
	                   // Default With the Logger and Recovery middleware already attached
					   
	r.GET("/ping", func(c *gin.Context) {

		var p Person // Person Object

		p.Name = "Sunil Kumar";
		p.Age = 24;

		c.Bind(&p)
         
		c.JSON(200, gin.H{
			"name": p.Name,
			"Age":p.Age,
		})
	})
	r.Run() // By default it serves on :8080 unless a
	        // PORT environment variable was defined.
}