package Tavern

import (
	"ddd-go/domain/product"
	"ddd-go/services/Order"
	"github.com/google/uuid"
	"testing"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}

	wine, err := product.NewProduct("Wine", "Healthy Beverage", 0.99)
	if err != nil {
		t.Error(err)
	}

	peanuts, err := product.NewProduct("Peanuts", "Healthy snacks", 0.99)
	if err != nil {
		t.Error(err)
	}

	products := []product.Product{
		beer, wine, peanuts,
	}

	return products

}

func TestTavern(t *testing.T) {
	products := init_products(t)

	os, err := Order.NewOrderService(Order.WithMemoryCustomerRepository(),
		Order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}

}

func Test_MongoTavern(t *testing.T) {
	products := init_products(t)

	os, err := Order.NewOrderService(
		Order.WithMongoCustomerRepository("mongodb://localhost:27017"),
		Order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}
	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}
}
