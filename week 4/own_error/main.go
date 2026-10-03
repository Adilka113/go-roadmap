package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type ValidationError struct {
	Field string
}

func (s ValidationError) Error() string {
	return fmt.Sprintf("пустое поле %s", s.Field)
}

func ValidateUser(u User) error {
	if len(u.Name) == 0 {
		return ValidationError{"Name"}
	}
	if u.Age == 0 {
		return ValidationError{"Age"}
	}
	return nil
}

func main() {
	Adil := User{Age: 20}
	err := ValidateUser(Adil)

	if err != nil {
		fmt.Println("Ошибка валидации:", err)
		return
	}
	fmt.Println("Пользователь инициализирован!")
}
