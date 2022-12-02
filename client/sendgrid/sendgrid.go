package sendgrid

import (
	"time"

	"github.com/kevinmidboe/planetposen-mail/client"
	"github.com/kevinmidboe/planetposen-mail/config"
)

// Client holds the HTTP client and endpoint information.
type Client struct {
	Endpoint   string
	APIKey     string
	HTTPClient client.HTTPClient
}

// Init sets up a new sendgrid client.
func (c *Client) Init(config *config.Config) error {
	timeout := 5 * time.Second
	c.Endpoint = config.SendGridAPIEndpoint
	c.APIKey = config.SendGridAPIKey
	c.HTTPClient = client.NewHTTPClient(client.Parameters{Timeout: &timeout})
	return nil
}

type sendEmailPayload struct {
	Personalizations []personalization `json:"personalizations"`
	From             email             `json:"from"`
	Content          []content         `json:"content"`
}

type personalization struct {
	To      []email `json:"to"`
	Subject string  `json:"subject"`
}

type email struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
