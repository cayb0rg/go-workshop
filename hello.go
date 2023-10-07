package main

import (
	"example/user/hello/models"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, world.")

	var characters = make([]models.Character, 0, 1)

	characters = append(characters,
		models.Character{
			Name:  models.Name{FirstName: "Gilfred", LastName: "Widowmaker"},
			Class: models.Class{ClassName: "Bard"},
			Race:  "Lizardfolk",
		},
	)

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	server.GET("/character", func(c *gin.Context) {
		firstName := c.Query("firstName")
		lastName := c.Query("lastName")

		if firstName == "" || lastName == "" {
			c.String(http.StatusNotFound, "Our server rolled a 1 on investigation :(")
			return
		}

		for _, thisCharacter := range characters {
			if thisCharacter.Name.FirstName == firstName || thisCharacter.Name.LastName == lastName {
				c.JSON(http.StatusOK, gin.H{
					"characters": thisCharacter,
				})
				return
			}
		}

		c.String(http.StatusNotFound, "Our server rolled a 1 on investigation :(")
	})

	server.Run("localhost:3000")
}
