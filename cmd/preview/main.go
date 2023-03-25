package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Product struct {
	ProductNo   int
	Name        string
	Image       string
	Quantity    int
	Price       float32
	Currency    string
}

type ReceiptPageData struct {
	PageTitle string
	OrderId   string
	Products  []Product
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

type Customer struct {
	FirstName     string
	LastName      string
	StreetAddress string
	ZipCode       string
	City          string
}

func serveTemplate(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("mail-templates/order-confirmation.html"))
	// tmpl := template.Must(w)
	data := EmailTemplateData{
		PageTitle: "Takk for din bestilling!",
		Site: "https://planet.schleppe.cloud",
		Date: time.Now().Format("2006-01-02"),
		Sum: 266.43,
		OrderId: "0upJLUYPEYaOCeQMxPc=",
		Products: []Product{
			{ProductNo: 1, Name: "Forrest", Image: "https://storage.googleapis.com/planetposen-images/838074447f08f03c4b75ac2030dcd01201c0656c.jpg", Quantity: 4, Price: 49.99, Currency: "NOK"},
			{ProductNo: 2, Name: "Cookie-Man Forrest", Image: "https://storage.googleapis.com/planetposen-images/2c47ed96b5e061d85f688849b998aa5e76c55c2a.jpg", Quantity: 3, Price: 99, Currency: "NOK"},
			{ProductNo: 3, Name: "Floral", Image: "https://planet.schleppe.cloud/email/items/item-3.jpg", Quantity: 1, Price: 129, Currency: "NOK"},
		},
		Customer: Customer{
			FirstName: "kevin",
			LastName: "Midbøe",
			StreetAddress: "Schleppegrells gate 18",
			ZipCode: "0001",
			City: "Oslo",
		},
	}

	b := &bytes.Buffer{}
	template := bufio.NewWriter(b)
	err := tmpl.Execute(template, data)
	if err != nil {
		fmt.Println(err)
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
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
