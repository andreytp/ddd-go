package memory

import (
	"ddd-go/domain/product"
	"github.com/google/uuid"
	"sync"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (m MemoryProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product
	for _, product := range products {
		products = append(products, product)
	}
	return products, nil
}

func (m MemoryProductRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, ok := m.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return product.Product{}, product.ErrProductNotFound
}

func (m MemoryProductRepository) Add(newProduct product.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[newProduct.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}
	m.products[newProduct.GetID()] = newProduct

	return nil
}

func (m MemoryProductRepository) Update(upProduct product.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[upProduct.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	m.products[upProduct.GetID()] = upProduct
	return nil
}

func (m MemoryProductRepository) Delete(id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(m.products, id)
	return nil
}
