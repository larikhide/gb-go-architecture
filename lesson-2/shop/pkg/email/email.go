package email

import (
	"net/smtp"
	"shop/models"
)

type EmailClient interface {
	SendOrderConfirmation(order *models.Order) error
}

type emailClient struct {
	cli *smtp.Client
}

func (email *emailClient) SendOrderConfirmation(order *models.Order) error {
https: //riptutorial.com/ru/go/example/20761/отправка-электронной-почты-с-помощью-smtp-sendmail---
}
