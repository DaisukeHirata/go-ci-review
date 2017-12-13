package main

import "testing"

func TestAddAndGetProductsInCart (t *testing.T) {
	c := New()
	c.Add("Apple")
//	c.Add("Grape")

	products := c.GetAll()
	if len(products) !=2 {
		t.Fatalf("The number of products is not correct.（Num of products : %d）", len(products))
	}
	if products[0] != "Apple" && products[1] != "Apple" {
		t.Error("Apple is not in the cart.")
		t.Log("Contents of the cart:", products)
	}
	if products[0] != "Orange" && products[1] != "Orange" {
		t.Error("Orange is not in the cart.")
		t.Log("Contents of the cart:", products)
	}
}
