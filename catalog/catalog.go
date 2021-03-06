package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
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

	return products.Products
}


func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadProducts()

	t := template.Must(template.ParseFiles("templates/catalog.html"))
	t.Execute(w, products)
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Enviando requisição para microsserviço de produto
	response, err := http.Get(baseUrlProducts + "/products/" + vars["id"])
	if err != nil {
		fmt.Printf("Falha ao carregar requisição HTTP %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("templates/view.html"))
	t.Execute(w, product)
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ListProducts)
	r.HandleFunc("/products/{id}", ShowProduct)
	http.ListenAndServe(":8083", r)
}
