// @title FizzBuzz implemention in Go
// @version 1.0
// @description This is a implemention of FizzBuzz in Golang. Github: ...

// @contact.name API Support
// @contact.email nourzad.shahrzad@gmail

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

package main

import (
	"FizzBuzz/handlers"
	"log"
	"net/http"

	_ "FizzBuzz/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	router := gin.Default()
	router.POST("/", POSTFUNC)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run()
	if err != nil {
		log.Fatalln("HTTP server error:", err)

	}
}

// @Summary return a fizz, buzz, fizzbuzz or the input number based on fizzbuzz game
// @ID fizzbuzz
// @Produce plain
// @Param inputNumber body string true "a number to check if its fizz, buzz, fizzbuzz or none"
// @Success 200 {object} string
// @Router / [post]
func POSTFUNC(c *gin.Context) {

	var reqBody int
	// log.Println("contenttype in binding", c.ContentType())
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Println("request binding error ", err)
		var errMsg string
		if reqBody == 0 {
			errMsg = "FizzBuzz Only accept normal numbers"
		}

		c.String(http.StatusBadRequest, "%s", errMsg+"\n")
		return

	}

	result, err := handlers.FizzBuzz(reqBody)

	if err != nil {
		errMsg := "Could not handel FizzBuzz\n"
		log.Println(errMsg, err)
		c.String(http.StatusInternalServerError, "%s", errMsg)
		return

	}

	c.String(http.StatusOK, "%s", result)
}
