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
	"fmt"

	_ "FizzBuzz/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	router := gin.Default()
	router.POST("/", POSTFUNC)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}

// @Summary return a fizz, buzz, fizzbuzz or the input number based on fizzbuzz game
// @ID fizzbuzz
// @Produce plain
// @Param inputNumber body string true "a number to check if its fizz, buzz, fizzbuzz or none"
// @Success 200 {object} string
// @Router / [post]
func POSTFUNC(c *gin.Context) {
	var con int

	if err := c.ShouldBindJSON(&con); err != nil {
		//fmt.Println("ERR in shouldbind", err)
		panic(err)
	}
	res := handlers.FizzBuzz(&con)
	fmt.Println(res)
	c.String(200, "%s", res)

}
