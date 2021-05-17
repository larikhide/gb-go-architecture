package email

import (
	"fmt"
	"net/smtp"
	"shop/models"
)

type EmailClient interface {
	SendOrderConfirmation(order *models.Order) error
}

type emailClient struct {
	cli      *smtp.Client
	username string
	password string
	host     string
}

func NewSMTPClient(host, username, password string) (*emailClient, error) {
	return &emailClient{
		host:     host,
		username: username,
		password: password,
	}, nil
}

func (email *emailClient) SendOrderConfirmation(order *models.Order) error {

	text := fmt.Sprintf("new order %d\nemail: %s\nphone: %s", order.ID, order.Email, order.Phone)

	auth := smtp.PlainAuth("", email.username, email.password, email.host)

	err := smtp.SendMail(email.host+":587", auth, email.username, []string{order.Email}, []byte(text))
	if err != nil {
		return err
	}
	return nil
}
