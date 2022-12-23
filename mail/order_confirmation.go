package mail

import (
	"context"
	"fmt"
	"time"
)

type OrderMailSender interface {
	SendOrderConfirmation(ctx context.Context, record Record) error
}

type Product struct {
	ProductNo   int
	Name        string
	Image       string
	Quantity    int
	Price       float32
	Currency    string
}

type Customer struct {
	FirstName     string
	LastName      string
	StreetAddress string
	ZipCode       string
	City          string
}

type OrderConfirmationData struct {
	// PageTitle string
	Email    string
	OrderId  string
	Customer Customer
	Products []Product
	Sum       float32
}

type EmailTemplateData struct {
	PageTitle string
	Site      string
	Date      string
	OrderId   string
	Customer  Customer
	Products  []Product
	Sum       float32
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
	emailTemplate.PageTitle = "Ordrebekreftelse fra planetposen.no"
	emailTemplate.Site = "https://planet.schleppe.cloud"
	emailTemplate.Date = time.Now().Format("2006-01-02")
	emailTemplate.Sum = payload.Sum
	emailTemplate.OrderId = payload.OrderId
	emailTemplate.Products = payload.Products
	emailTemplate.Customer = payload.Customer

	orderConfirmationEmailData := buildOrderConfirmation(emailTemplate)
	if orderConfirmationEmailData == nil {
		return nil, fmt.Errorf("couldn't build order confirmation template for orderId %s", payload.OrderId)
	}

	orderConfirmationEmailData.ToEmail = payload.Email

	return orderConfirmationEmailData, nil
}
