package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// Set up routes
	// Home route
	http.HandleFunc("/", homecontroller.Welcome)

	// Categories routes
	http.HandleFunc("/categories", categorycontroller.IndexCategories)
	http.HandleFunc("/categories/create", categorycontroller.CreateCategory)
	http.HandleFunc("/categories/edit", categorycontroller.EditCategory)
	http.HandleFunc("/categories/delete", categorycontroller.DeleteCategory)
	
	// Products routes
	http.HandleFunc("/products", productcontroller.IndexProducts)
	http.HandleFunc("/products/show", productcontroller.ShowProduct)
	http.HandleFunc("/products/create", productcontroller.CreateProduct)
	http.HandleFunc("/products/edit", productcontroller.EditProduct)
	http.HandleFunc("/products/delete", productcontroller.DeleteProduct)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
