package productcontroller

import (
	"go-web-native/models/productmodel"
	"html/template"
	"net/http"
	"strings"
)

func IndexProducts(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	// println(products[0].Category.Name) // Debugging line to check category name
	data := map[string]any{
		"products": products,
	}
	// Handle the products page request
	temp, err := template.ParseFiles("views/products/products.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
	// Handle the form submission for creating a product
}
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Handle the form submission for creating a product
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	// Handle the form submission for editing a product
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Handle the form submission for deleting a product
}

func limitWords(s string, max int) string {
	words := strings.Fields(s)
	if len(words) <= max {
		return s
	}
	return strings.Join(words[:max], " ")
}