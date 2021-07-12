package main

import "fmt"

type User struct {
	Name  string
	Email string
}

type UserNotifier interface {
	SendMsg(user *User, msg string) error
}

type EmailNotifier struct {
	//for API
}

func (notifier EmailNotifier) SendMsg(user *User, msg string) error {
	_, err := fmt.Printf("Sending mail to %s with content %s\n", user.Name, msg)
	return err
}

func main() {
	user := User{"Shivesh", "shivesh.ojha@gmail.com"}
	fmt.Printf("Welcome %s\n", user.Name)

	var notify UserNotifier
	notify = EmailNotifier{}
	notify.SendMsg(&user, "Interfaces are used in the backend!")
}
