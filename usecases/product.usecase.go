package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecases struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecases {
	return ProductUsecases{
		repository: repo,
	}
}

func (pu *ProductUsecases) GetProducts() ([]model.Product, error) {
	products, err := pu.repository.GetProducts()
	return products, err
}

func (pu *ProductUsecases) GetProductById(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id)
	return product, err
}

func (pu *ProductUsecases) UpdateProductById(id int, p model.Product) (*model.Product, error) {
	product, err := pu.repository.UpdateProductById(id, p)
	return product, err
}

func (pu *ProductUsecases) CreateProduct(p model.Product) (model.Product, error) {
	id, err := pu.repository.CreateProduct(p)
	if err != nil {
		return model.Product{}, err
	}

	p.ID = id
	return p, nil
}
