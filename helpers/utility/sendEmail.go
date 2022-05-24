package utility

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"

	"Hybrid-blog/helpers/config"
)

var (
	emailFrom = ""
	password  = ""
)

func SendEmail(emailTo, message string) {
	loadInfoEmail()
	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", "Reset Password")
	m.SetBody("text/plain", "Link: "+message)
	d := gomail.NewDialer("smtp.gmail.com", 587, emailFrom, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}

func loadInfoEmail() {
	env := config.GetEnvValue()
	emailFrom = env.Gmail.User
	password = env.Gmail.Password
}
