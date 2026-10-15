package main

import (
	"errors"
	"fmt"
)

var ErrDivByZero = errors.New("divide by zero forbidden")

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

func main() {
	a := 10.0
	b := 0.0
	res, err := Divide(a, b)

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println(res)
}
