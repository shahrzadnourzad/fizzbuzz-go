package handlers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
)

func FizzBuzz(c *int) string {

	con := *c
	if con > 0 && con < 101 {

		if con%3 == 0 && con%5 == 0 {
			return `"FizzBuzz"` + "\n"

		} else if con%3 == 0 {
			return `"Fizz"` + "\n"

		} else if con%5 == 0 {
			return `"Buzz"` + "\n"

		} else {
			return `"` + strconv.Itoa(con) + `"` + "\n"
		}

	} else {
		return outsideIntervalFizzBuzz(&con)

	}
}

func outsideIntervalFizzBuzz(c *int) string {

	con := *c
	data := []byte(strconv.Itoa(con))
	res, err := http.Post("https://fizzbuzz-oracle.b17g.services/predict", "text/html; charset=UTF-8", bytes.NewBuffer(data))

	if err != nil {
		
		panic(err)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
