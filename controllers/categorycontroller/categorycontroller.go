package categorycontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func IndexCategories(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}
	// Handle the categories page request
	temp, err := template.ParseFiles("views/categories/categories.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// Handle the categories page request
		temp, err := template.ParseFiles("views/categories/create_category.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		// Handle the form submission for creating a category
		var category entities.Category
		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Create(category); !ok {
			template, err := template.ParseFiles("views/categories/create_category.html")
			if err != nil {
				http.Error(w, "Error loading template", http.StatusInternalServerError)
				return
			}
			template.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
		// http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func EditCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Handle the categories page request
		temp, err := template.ParseFiles("views/categories/edit_category.html")
		if err != nil {
			panic(err)
		}
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

		category := categorymodel.GetById(id)
		if category == nil {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}

		data := map[string]any{
			"Id":   category.Id,
			"Name": category.Name,
		}
		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var category entities.Category
		idString := r.FormValue("id")
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
		category.Name = r.FormValue("name")
		category.Id = uint(id)
		category.UpdatedAt = time.Now()
		if ok := categorymodel.Update(category); !ok {
			temp, err := template.ParseFiles("views/categories/edit_category.html")
			if err != nil {
				http.Error(w, "Error loading template", http.StatusInternalServerError)
				return
			}
			temp.Execute(w, nil)
			return // Hentikan di sini agar tidak lanjut ke Redirect
		}

		// Redirect hanya dijalankan kalau update sukses
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
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
		categorymodel.Delete(id)
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}
