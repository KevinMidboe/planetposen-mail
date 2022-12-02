package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Product struct {
	Name        string
	Image       string
	Description string
	Quantity    int
	Price       float32
	Currency    string
}

type ReceiptPageData struct {
	PageTitle string
	OrderId   string
	Products  []Product
}

func serveTemplate(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("mail/mail-template_order-confirmation.html"))
	// tmpl := template.Must(w)
	data := ReceiptPageData{
		PageTitle: "Planetposen purchase",
		OrderId:   "fb9a5910-0dcf-4c65-9c25-3fb3eb883ce5",
		Products: []Product{
			{Name: "Forrest", Image: "https://planet.schleppe.cloud/email/items/item-1.jpg", Description: "Sneaker Maker", Quantity: 4, Price: 49.99, Currency: "NOK"},
			{Name: "Cookie-Man Forrest", Image: "https://planet.schleppe.cloud/email/items/item-2.jpg", Description: "Boots Brothers", Quantity: 3, Price: 99, Currency: "NOK"},
			{Name: "Floral", Image: "https://planet.schleppe.cloud/email/items/item-3.jpg", Description: "Swiss Made", Quantity: 1, Price: 129, Currency: "NOK"},
		},
	}

	var b bytes.Buffer
	template := bufio.NewWriter(&b)
	err := tmpl.Execute(template, data)
	if err != nil {
		fmt.Println(err)
	}

	response.Header().Set("Content-Type", "text/html")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(b.String()))
}

func main() {
	ADDRESS := ":5006"
	fmt.Printf("Serving preview of template at %s\n", ADDRESS)
	http.HandleFunc("/", serveTemplate)
	log.Fatal(http.ListenAndServe(ADDRESS, nil))
}
