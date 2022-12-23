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
	tmpl := template.Must(template.ParseFiles("mail-templates/order-confirmation.html"))
	// tmpl := template.Must(w)
	data := ReceiptPageData{
		PageTitle: "Ordrebekreftelse planetposen",
		OrderId:   "k6z6wq2J_pFMBY78Sp0=",
		Products: []Product{
			{Name: "Forrest", Image: "https://storage.googleapis.com/planetposen-images/838074447f08f03c4b75ac2030dcd01201c0656c.jpg", Description: "Sneaker Maker", Quantity: 4, Price: 49.99, Currency: "NOK"},
			{Name: "Cookie-Man Forrest", Image: "https://storage.googleapis.com/planetposen-images/2c47ed96b5e061d85f688849b998aa5e76c55c2a.jpg", Description: "Boots Brothers", Quantity: 3, Price: 99, Currency: "NOK"},
			{Name: "Floral", Image: "https://planet.schleppe.cloud/email/items/item-3.jpg", Description: "Swiss Made", Quantity: 1, Price: 129, Currency: "NOK"},
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
