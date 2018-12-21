package routes

import (
	"github.com/gorilla/mux"
	"git.oxy.sh/oxy/zeal/models"
	"git.oxy.sh/oxy/zeal/render"
	"net/http"
	"strconv"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	render.JSON(w, http.StatusOK, products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	i, err := strconv.Atoi(id)

	if err != nil {
		render.JSONErr(w, http.StatusBadRequest, "ID must be numeric")
		return
	}

	product := models.GetProductById(i)

	if product == nil {
		render.JSONErr(w, http.StatusNotFound, "Resource not found")
		return
	}

	render.JSON(w, http.StatusOK, product)
}