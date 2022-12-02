package mail

import (
	"html/template"
	"strings"
)

func buildOrderConfirmation(templateData EmailTemplateData) *OrderConfirmationEmailData {
	subject := "Orderbekreftelse fra planetposen.no"

	data := &OrderConfirmationEmailData{
		Subject:   subject,
		FromName:  "noreply@kevm.dev",
		FromEmail: "noreply@kevm.dev",
	}

	tmpl := template.Must(template.ParseFiles("mail/mail-template_order-confirmation.html"))
	b := new(strings.Builder)
	err := tmpl.Execute(b, templateData)
	if err != nil {
		return nil
	}

	data.Markup = b.String()
	return data
}
