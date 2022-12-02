package sendgrid

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kevinmidboe/planetposen-mail/client"
	"github.com/kevinmidboe/planetposen-mail/mail"
	"net/http"
)

// SendOrderConfirmation sends an order confirmation.
func (c *Client) SendOrderConfirmation(ctx context.Context, record mail.OrderConfirmationEmailData) error {
	reqBody := sendEmailPayload{
		Personalizations: []personalization{
			{
				To: []email{
					{
						Email: record.ToEmail,
					},
				},
				Subject: record.Subject,
			},
		},
		From: email{
			Email: record.FromEmail,
			Name:  record.FromName,
		},
		Content: []content{
			{
				Type:  "text/html",
				Value: record.Markup,
			},
		},
	}
	jsonPayload, err := json.Marshal(reqBody)

	if err != nil {
		return fmt.Errorf("error marshalling sendEmailPayload: %w", err)
	}
	reqData := client.HTTPRequestData{
		Method: http.MethodPost,
		URL:    fmt.Sprintf("%s/v3/mail/send", c.Endpoint),
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", c.APIKey),
		},
		PostPayload: jsonPayload,
	}
	_, err = c.HTTPClient.RequestBytes(ctx, reqData)
	if err != nil {
		return fmt.Errorf("error making request to sendgrid to send email: %w", err)
	}

	return nil
}
