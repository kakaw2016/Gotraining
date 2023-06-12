package main

import (
	"fmt"
	"net/smtp"
)

const (
	username = "your_username"
	password = "your_password"
	host     = "smtp.example.com"
	port     = "587"
)

func sendMail(to, subject, body string) error {
	from := username

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", username, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	to := "recipient@example.com"
	subject := "Notification"
	body := "This is a test email notification."

	err := sendMail(to, subject, body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email sent successfully.")
	}
}
