package models

import "testing"

func TestGetProductById(t *testing.T) {
	product := GetProductById(1)

	t.Logf("Product: %+v", product)

	if product.Code != "TestingProduct" {
		t.Errorf("Couldn't find product based on ID.")
	}
}
