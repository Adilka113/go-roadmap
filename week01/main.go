package main

import "fmt"

func main() {
	// 1
	var a float64
	var b float64

	var sum float64

	fmt.Println("Введите число а:")
	fmt.Scan(&a)

	fmt.Println("Введите число b:")
	fmt.Scan(&b)

	sum = a + b

	fmt.Println("Итоговый ответ:", sum)

	//2
	var c int

	fmt.Print("Введите число:")
	fmt.Scan(&c)

	if c%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}

	//3
	d := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	e := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(d[i], "x", e[j], "=", d[i]*e[j])
		}
	}
}
