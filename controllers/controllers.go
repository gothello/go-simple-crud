package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gothello/go-web-studies/models"
)

var templates = template.Must(template.ParseGlob("./templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	templates.ExecuteTemplate(w, "Index", products)
	//json, _ := json.Marshal(products)
	//	w.Write(json)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")
		id := r.FormValue("id")

		priceOk, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Failed to convert price or float", err)
		}

		idOK, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Failed to convert id to int", err)
		}

		quantityOk, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Println("Failed to convert quantity to integer", err)
		}

		models.SaveProduct(idOK, name, description, priceOk, quantityOk)
	}

	http.Redirect(w, r, "/", 302)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idOK, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Failed to convert id to int")
	}

	product := models.GetProduct(idOK)

	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idOK, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Failed to convert to int")
		}

		priceOK, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Failed to convert price to float64")
		}

		quantityOK, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Failed to convert quantity to int")
		}

		models.UpdateProduct(
			idOK,
			name,
			description,
			priceOK,
			quantityOK,
		)
	}

	http.Redirect(w, r, "/", 302)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no cache, no store, must-revalidate, private, max-age=0")
	w.Header().Set("Pragma", "no-cahche")
	w.Header().Set("X-Accel-Expires", "0")
	w.Header().Set("Expires", "0")

	id := r.URL.Query().Get("id")

	idOK, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Failed to convert id to integer", err)
	}

	models.DeleteProduct(idOK)

	http.Redirect(w, r, "/", 302)
}
