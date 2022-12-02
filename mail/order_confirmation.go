package mail

import (
	"context"
	"fmt"
)

type OrderMailSender interface {
	SendOrderConfirmation(ctx context.Context, record Record) error
}

type Product struct {
	Name        string
	Image       string
	Description string
	Quantity    int
	Price       float32
	Currency    string
}

type OrderConfirmationData struct {
	// PageTitle string
	Email    string
	OrderId  string
	Products []Product
}

type EmailTemplateData struct {
	PageTitle string
	OrderId   string
	Products  []Product
}

type OrderConfirmationEmailData struct {
	Subject   string
	FromName  string
	FromEmail string
	ToEmail   string
	Markup    string
}

type Record struct {
	Email string
	// FullName									 string
	Status                     string
	OrderConfirmationEmailData OrderConfirmationEmailData
}

func OrderConfirmation(payload OrderConfirmationData) (*OrderConfirmationEmailData, error) {
	var emailTemplate EmailTemplateData
	emailTemplate.PageTitle = "Planetposen purchase"
	emailTemplate.OrderId = payload.OrderId
	emailTemplate.Products = payload.Products

	orderConfirmationEmailData := buildOrderConfirmation(emailTemplate)
	if orderConfirmationEmailData == nil {
		return nil, fmt.Errorf("couldn't build order confirmation template for orderId %s", payload.OrderId)
	}

	orderConfirmationEmailData.ToEmail = payload.Email

	return orderConfirmationEmailData, nil
}
