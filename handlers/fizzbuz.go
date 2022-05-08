package handlers

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func FizzBuzz(reqBody int) (string, error) {

	if reqBody >= 1 && reqBody <= 100 {

		if reqBody%3 == 0 && reqBody%5 == 0 {
			return `"FizzBuzz"` + "\n", nil

		} else if reqBody%3 == 0 {
			return `"Fizz"` + "\n", nil

		} else if reqBody%5 == 0 {
			return `"Buzz"` + "\n", nil

		} else {
			return `"` + strconv.Itoa(reqBody) + `"` + "\n", nil
		}

	} else {
		resp, err := outsideIntervalFizzBuzz(reqBody)
		return resp, err
	}
}

func outsideIntervalFizzBuzz(reqBody int) (string, error) {

	data := (strconv.Itoa(reqBody))
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Post("https://fizzbuzz-oracle.b17g.services/predict", "text/plain; charset=UTF-8", strings.NewReader(data))
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("received" + strconv.Itoa(resp.StatusCode) + "from Oracle")
	}

	bytes, err := io.ReadAll(resp.Body)
	if err!= nil{
		return "",errors.New("can not read the oracle response body")
	}
	respBody := string(bytes)
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/plain; charset=utf-8" {
		return "", errors.New("the respponse is not expected type")

	}

	if !eval(respBody, data) {
		return "", errors.New("oracle response is not valid")
	}
	return respBody, nil
}

func eval(resp string, data string) bool {

	expresp := []string{
		`"Fizz"` + "\n",
		`"Buzz"` + "\n",
		`"FizzBuzz"` + "\n",
		`"` + data + `"` + "\n",
	}

	for _, e := range expresp {
		if resp == e {
			return true
		}
	}

	return false
}
