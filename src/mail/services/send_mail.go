package services

import (
	"fmt"
	"time"
)

type Mail struct {
	To      string
	From    string
	Subject string
	Content string
}

func SendMail(mail Mail) {
	fmt.Println("mail sending started", mail)
	time.Sleep(time.Second * 10)
	fmt.Println("Background mail send complete", mail)
}
