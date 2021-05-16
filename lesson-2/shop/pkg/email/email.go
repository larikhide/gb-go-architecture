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
	cli *smtp.Client
}

func NewSMTPClient(network, host string) (*emailClient, error) {
	conn, err := net.Dial(network, host)
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

	//Q: Не уверен, что хорошо делать аунтификацию прямо здесь. На мой взгляд это нужно делать в мэйне.
	// Как все же правильнее?
	//auth := smtp.PlainAuth("", "someuser@example.com", "password", "mail.example.com")

	err := smtp.SendMail("mail.example.com:25", auth, "someuser@example.com", []string{order.Email}, []byte(text))
	if err != nil {
		return err
	}
	return nil
}
