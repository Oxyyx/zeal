package models

type Product struct {
	Cohort
	Code string `json:"code"`
	Price uint `json:"price"`
}

func GetAllProducts() *[]Product {
	products := []Product{}
	db.Find(&products)

	return &products
}

func GetProductById(id int) *Product {
	product := Product{}

	if db.Find(&product, id).RecordNotFound() {
		return nil
	}

	return &product
}