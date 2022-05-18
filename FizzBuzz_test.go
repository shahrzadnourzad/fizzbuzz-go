package main

import (
	"FizzBuzz/handlers"
	"fmt"
	"testing"
)

func TestFizz(t *testing.T) {

	expres := []struct {
		in   int
		want string
	}{
		{1, `"1"` + "\n"},
		{3, `"Fizz"` + "\n"},
		{5, `"Buzz"` + "\n"},
		{15, `"FizzBuzz"` + "\n"},
		{105, `"FizzBuzz"` + "\n"},
	}

	for _, num := range expres {
		funcRes, _ := handlers.FizzBuzz(num.in)
		fmt.Println("funcres:", funcRes)
		fmt.Println("num.wan:", num.want)
		if funcRes != num.want {
			t.Errorf("Fizzbuzz(%d): expected %s, but func response is %s", num.in, num.want, funcRes)
		}
	}
}
