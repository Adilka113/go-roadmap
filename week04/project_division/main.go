package main

import (
	"fmt"
	"project_division/mathutils"
)

func main() {
	a := 10.0
	b := 20.0

	fmt.Println(mathutils.Add(a, b), mathutils.Multiply(a, b))
}
