package smtp

import (
	"errors"
	"shop/models"
)

var (
	ErrIncorrectEmail = errors.New("email is incorrect")
)

type SmtpAPI interface {
	SendOrderNotification(order *models.Order) error
}

type smtpAPI struct {
}
