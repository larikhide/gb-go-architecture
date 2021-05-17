package email

import (
	"fmt"
	"net"
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

func NewSMTPClient(network string, host string) (*emailClient, error) {
	conn, err := net.Dial(network, host+":587")
	if err != nil {
		return nil, err
	}
	cli, err := smtp.NewClient(conn, host)
	if err != nil {
		return nil, err
	}
	return &emailClient{
		cli: cli,
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
