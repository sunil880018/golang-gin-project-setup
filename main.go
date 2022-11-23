package main

import "github.com/gin-gonic/gin"
import "net/http"

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	r := gin.Default() // Creates a gin router with default middleware:
	                   // Default With the Logger and Recovery middleware already attached
					   
	r.GET("/ping", func(c *gin.Context) { 

		var p Person // Person Object

		p.Name = "Sunil Kumar"; // Note that p.Name becomes "name" in the JSON
		p.Age = 24;

		c.Bind(&p)
         
		// c.JSON() return output in json format
		c.JSON(http.StatusOK, gin.H{ // gin.H is a shortcut for map[string]interface{}
			"name": p.Name,
			"Age":p.Age,
		})
	})
	r.Run() // By default it serves on :8080 unless a
	        // PORT environment variable was defined.
}