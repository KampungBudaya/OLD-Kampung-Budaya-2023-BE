package email

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"text/template"

	"gopkg.in/gomail.v2"
)

type Gomail struct {
	message *gomail.Message
	dialer  *gomail.Dialer
}

func NewMailClient() *Gomail {
	port, err := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	if err != nil {
		log.Fatalf("[tedxub2023-api-http] failed to convert smtp port: %s\n", err.Error())
	}

	return &Gomail{
		gomail.NewMessage(),
		gomail.NewDialer(
			os.Getenv("CONFIG_SMTP_HOST"),
			port,
			os.Getenv("CONFIG_AUTH_EMAIL"),
			os.Getenv("CONFIG_AUTH_PASSWORD"),
		)}
}

func (g *Gomail) SetSender(sender string) {
	g.message.SetHeader("From", sender)
}

func (g *Gomail) SetReciever(to ...string) {
	g.message.SetHeader("To", to...)
}

func (g *Gomail) SetSubject(subject string) {
	g.message.SetHeader("Subject", subject)
}

func (g *Gomail) SetBodyHTML(email, password string) error {
	var body bytes.Buffer
	path := "template.html"
	t, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	t.Execute(&body, struct {
		Email    string
		Password string
	}{
		Email:    email,
		Password: password,
	})
	g.message.SetBody("text/html", body.String())
	return nil
}

func (g *Gomail) SendMail() error {
	if err := g.dialer.DialAndSend(g.message); err != nil {
		return err
	}
	return nil
}
