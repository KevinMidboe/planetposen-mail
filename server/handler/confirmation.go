// Package handler contains handlers for events.
package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kevinmidboe/planetposen-mail/client/sendgrid"
	"github.com/kevinmidboe/planetposen-mail/mail"
	"net/http"
)

func SendOrderConfirmation(s *sendgrid.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload, err := getOrderConfirmationPayload(r)

		if err != nil {
			handleError(w, err, "unable to parse order payload", http.StatusBadRequest, true)
			return
		}
		mailData, err := mail.OrderConfirmation(*payload)

		err = s.SendOrderConfirmation(ctx, *mailData)
		if err != nil {
			fmt.Println(err)
			handleError(w, err, "error from sendgrid ", http.StatusInternalServerError, true)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		responseJSON, _ := json.Marshal(struct {
			Message   string `json:"message"`
			OrderId   string `json:"orderId"`
			Recipient string `json:"recipient"`
		}{
			Message:   "Successfully sent email",
			OrderId:   payload.OrderId,
			Recipient: payload.Email,
		})
		w.Write(responseJSON)
	}
}

func getOrderConfirmationPayload(r *http.Request) (*mail.OrderConfirmationData, error) {
	decoder := json.NewDecoder(r.Body)

	var payload mail.OrderConfirmationData
	err := decoder.Decode(&payload)
	if err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}

	return &payload, nil
}
