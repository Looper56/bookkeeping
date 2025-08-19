package email

import (
	"bookkeeping/config"

	"gopkg.in/gomail.v2"
)

func SendMail(from, subject, body string, mailTo []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(config.Config.Email.Account, from))
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	send := gomail.NewDialer(config.Config.Email.Host, config.Config.Email.Port,
		config.Config.Email.Account, config.Config.Email.Password)
	err := send.DialAndSend(m)
	return err
}
