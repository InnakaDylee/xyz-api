package util

import (
	"fmt"
	"log"
	"strconv"
	configs "xyz/packages/databases/config"

	"github.com/go-gomail/gomail"
)

func SendEmailNotification(email string, token string) error {
	config, err := configs.LoadConfig()
	if err != nil {
		return err
	}
	to := []string{email}

	m := gomail.NewMessage()
	m.SetHeader("From", config.SMTP.SMTP_USERNAME)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "Activation Email")
	m.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
		<html>
		<head>
		<title>Account Activation</title>
		</head>
		<body>
		<h1>Activate Your Account</h1>
		<p>Hello %s,</p>
		<p>Please click the button below to activate your account:</p>
		<a href="http://localhost:8000/api/v1/user/activate?token=%s" 
			style="display: inline-block; padding: 10px 20px; font-size: 16px; color: #fff; background-color: #007bff; border: none; border-radius: 5px; text-decoration: none;">
			Activate Account
		</a>
		<p>If you did not request this, please ignore this email.</p>
		</body>
		</html>
	`, email, token))

	Port, _ := strconv.Atoi(config.SMTP.SMTP_PORT)

	dialer := gomail.NewDialer(
		config.SMTP.SMTP_SERVER,
		Port,
		config.SMTP.SMTP_USERNAME,
		config.SMTP.SMTP_PASSWORD,
	)

	if err := dialer.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	return nil
}