package utils

import (
	"fmt"
	"net/smtp"
)

func SendVerificationEmail(email, token string) error {
	from := "no-reply@poc.com"
	password := "sua_senha"
	to := []string{email}
	smtpHost := "smtp.seuservidor.com"
	smtpPort := "587"

	msg := []byte("Subject: Verifique seu e-mail\r\n" +
		"\r\n" +
		"Clique para verificar: http://localhost:8080/verify-email/" + token + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)
	return smtp.SendMail(addr, auth, from, to, msg)
}
