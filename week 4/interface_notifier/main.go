package main

import "fmt"

type Notifier interface {
	Notify(m string)
}

type EmailNotifier struct{}
type TelegramNotifier struct{}

func (n EmailNotifier) Notify(m string) {
	fmt.Println("Отправлено письмо:", m)
}

func (n TelegramNotifier) Notify(m string) {
	fmt.Println("Отправлено сообщение в Telegram:", m)
}

func SendNotification(n Notifier, m string) {
	n.Notify(m)
}

func main() {
	email := EmailNotifier{}
	telegram := TelegramNotifier{}

	SendNotification(email, "Письмо")
	SendNotification(telegram, "Сообщение")

}
