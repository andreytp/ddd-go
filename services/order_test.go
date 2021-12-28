package services

import (
	"ddd-go/aggregate"
	"github.com/google/uuid"
	"testing"
)

func TestOrder_NewOrderService(t *testing.T) {
	products := initProducts(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}

func initProducts(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}

	wine, err := aggregate.NewProduct("Wine", "Healthy Beverage", 0.99)
	if err != nil {
		t.Error(err)
	}

	peanuts, err := aggregate.NewProduct("Peanuts", "Healthy snacks", 0.99)
	if err != nil {
		t.Error(err)
	}

	products := []aggregate.Product{
		beer, wine, peanuts,
	}

	return products

}
