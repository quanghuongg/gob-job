package main

import (
	"demo-job/config"
	"demo-job/queue"
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"time"
)

var emailAuth smtp.Auth

type Sender struct {
	Email string
}

func SendEmailSMTP(to []string) (bool, error) {
	emailHost := config.GetString("mail.host")
	emailFrom := config.GetString("mail.username")
	emailPassword := config.GetString("mail.password")
	emailPort :=  config.GetInt("mail.port")

	emailAuth = smtp.PlainAuth("", emailFrom, emailPassword, emailHost)
	emailBody := "Mail from golang bot"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + "Test Email" + "!\n"
	msg := []byte(subject + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s: %d", emailHost, emailPort)

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (s Sender) Process() {
	_, err := SendEmailSMTP(strings.Split(s.Email, ","))
	if err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("Sent :", s.Email)
	}
	time.Sleep(time.Second * 1)
}

func main() {
	fmt.Println("demo job")
	emails := []string{
		"huongnq4@gmail.com",
		"quanghuongitus@gmail.com",
		"huongnq3@gmail.com",
		"huongnq4@gmail.com",
		"huongnq5@gmail.com",
	}
	jobQueue := queue.NewJobQueue(2)
	jobQueue.Start()
	fmt.Println(len(emails))
	for i := 0; i < len(emails); i++ {
		sender := Sender{Email: emails[i]}
		jobQueue.Push(sender)
	}
	time.AfterFunc(time.Second*20, func() {
		jobQueue.Stop()
	})
	time.Sleep(time.Second * 6)

}
