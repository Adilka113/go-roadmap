package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

func (n User) IsAdult() bool {
	return n.Age >= 18
}

type Rectangle struct {
	Width  int
	Height int
}

func (n Rectangle) Area() int {
	return n.Height * n.Width
}

type Address struct {
	City, Street string
}

type User1 struct {
	Name  string
	Age   int
	Email string
	Address
}

func main() {
	petya := User{"Петя", 18, "petya@gmail.com"}

	fmt.Printf("%+v", petya)
	fmt.Println(petya.IsAdult())

	r := Rectangle{10, 15}

	fmt.Println(r.Area())

	vasya := User1{"Вася", 18, "vasya@gmail.com", Address{"Moscow", "Arbat"}}

	fmt.Println("Имя:", vasya.Name+", Город:", vasya.City+", Улица:"+vasya.Street)
}
