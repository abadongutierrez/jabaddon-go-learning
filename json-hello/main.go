package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	ProductID      int    `json:"productId"`
	Manufaturer    string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

func main() {
	product := &Product{
		ProductID:      123,
		Manufaturer:    "Big Box Company",
		PricePerUnit:   "12.99",
		Sku:            "34AS1sy",
		Upc:            "12314123123",
		QuantityOnHand: 28,
		ProductName:    "Gizmo",
	}
	productJson, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(productJson))

	jsonProduct := `{"productId":123,
		"manufacturer":"Big Box Company",
		"sku":"34AS1sy",
		"upc":"12314123123",
		"pricePerUnit":"12.99",
		"quantityOnHand":28,
		"productName":"Gizmo"
	}`
	newProduct := Product{}
	newErr := json.Unmarshal([]byte(jsonProduct), &newProduct)
	if newErr != nil {
		log.Fatal(newErr)
	}
	fmt.Printf("json unmarshal product: %s", newProduct.ProductName)
}
