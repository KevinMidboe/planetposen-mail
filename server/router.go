package server

import (
	"github.com/kevinmidboe/planetposen-mail/server/handler"
)

const v1API string = "/api/v1/"

func (s *Server) setupRoutes() {
	s.Router.HandleFunc("/_healthz", handler.Healthz).Methods("GET").Name("Health")

	api := s.Router.PathPrefix(v1API).Subrouter()
	api.HandleFunc("/send-confirmation", handler.SendOrderConfirmation(s.SendGrid)).Methods("POST").Name("SendOrderConfirmation")
}
