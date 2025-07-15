package productcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"go-web-native/models/productmodel"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
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
	categories := categorymodel.GetAll()
	dataCategory := map[string]any{
		"categories": categories,
	}

	if r.Method == "GET" {
		// Handle the products page request
		temp, err := template.ParseFiles("views/products/create_product.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, dataCategory)
	}

	if r.Method == "POST" {
		// Handle the form submission for creating a category
		var product entities.Product

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil || stock < 0 {
			http.Error(w, "Invalid stock value", http.StatusBadRequest)
			return
		}

		product.Name = r.FormValue("name")
		product.Description = r.FormValue("description")
		product.Stock = int64(stock)
		product.Category.Id = uint(categoryId)
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if ok := productmodel.Create(product); !ok {
			template, err := template.ParseFiles("views/categories/create_category.html")
			if err != nil {
				http.Error(w, "Error loading template", http.StatusInternalServerError)
				return
			}
			template.Execute(w, nil)
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
		// http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	// Handle the form submission for editing a product
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Handle the form submission for deleting a product
		if r.Method == "GET" {
		idString := r.URL.Query().Get("id")
		if idString == "" {
			http.Error(w, "Category ID is required", http.StatusBadRequest)
			return
		}
		strconv.Atoi(idString)
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}
		productmodel.Delete(id)
		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func limitWords(s string, max int) string {
	words := strings.Fields(s)
	if len(words) <= max {
		return s
	}
	return strings.Join(words[:max], " ")
}
