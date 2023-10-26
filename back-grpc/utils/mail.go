package utils

import "net/smtp"

func SendMail(cfg *Config, to []string, subject string, body string) error {

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: " + subject + "\r\n" + "\r\n" + body + "\r\n")

	auth := smtp.PlainAuth("", cfg.EmailUser, cfg.EmailPass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, cfg.EmailUser, to, message)
	if err != nil {
		return err
	}

	return nil

}
