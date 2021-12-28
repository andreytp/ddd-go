package memory

import (
	"ddd-go/aggregate"
	"ddd-go/domain/product"
	"testing"

	"github.com/google/uuid"
)

func TestMemoryProductRepository_Add(t *testing.T) {
	repo := New()
	newProduct, err := aggregate.NewProduct("Beer", "Good for your health", 1.99)
	if err != nil {
		t.Error(err)
	}
	repo.Add(newProduct)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}
}

func TestMemoryProductRepository_Get(t *testing.T) {
	repo := New()
	existingProduct, err := aggregate.NewProduct("Beer", "Good for your health", 1.99)
	if err != nil {
		t.Error(err)
	}
	repo.Add(existingProduct)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get product by id",
			id:          existingProduct.GetID(),
			expectedErr: nil,
		}, {
			name:        "Get non existing product by id",
			id:          uuid.New(),
			expectedErr: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})

	}
}

func TestMemoryProductRepository_Delete(t *testing.T) {
	repo := New()
	existingProduct, err := aggregate.NewProduct("Beer", "Good for your health", 1.99)
	if err != nil {
		t.Error(err)
	}
	repo.Add(existingProduct)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}

	err = repo.Delete(existingProduct.GetID())
	if err != nil {
		t.Error(err)
	}
	if len(repo.products) != 0 {
		t.Errorf("Expected 0 products, got %d", len(repo.products))
	}
}