package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)


type Product struct {
	Uuid string `json:"uuid"`
	Product string `json:"product"`
	Price float64 `json:"price,string"`
}


type Products struct {
	Products []Product
}


var baseUrlProducts string


func init() {
	// Carregando arquivo .env da pasta product
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Falha ao carregar arquivo .env")
	}

	// Recebe o valor atríbuido a variável PRODUCT_URL=VALOR no arquivo .env 
	baseUrlProducts = os.Getenv("PRODUCT_URL")
}

func loadProducts() []Product {
	// Enviando requisição para microsserviço de produto
	response, err := http.Get(baseUrlProducts + "/products")
	if err != nil {
		fmt.Println("Erro de HTTP")
	}

	data, _ := ioutil.ReadAll(response.Body)

	var products Products
	json.Unmarshal(data, &products)
	fmt.Println(string(data))

	return products.Products
}

func main() {
	loadProducts()
}
