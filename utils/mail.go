package utils

import (
	"fmt"
	"go-core/databases/models"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

var EmailChannel = make(chan *models.User)

func SendMails() {
	for user := range EmailChannel {
		SendMail(user)
	}
}

func SendMail(user *models.User) {
	err := godotenv.Load()
	if err != nil {
		LogError(err)
		return
	}

	from := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	// Receiver email address
	to := []string{
		"tvhoan.2908@gmail.com",
	}

	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message
	msg := "Subject: Tài khoản mới trên website \n" +
		fmt.Sprintf("Tài khoản %s vừa mới được tạo.", user.Username)
	message := []byte(msg)

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		LogError(err)
		return
	}

	fmt.Println("Email sent successfully !")
}
