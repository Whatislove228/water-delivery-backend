package service

import (
	"context"
	"fmt"

	"water-delivery/internal/domain"
	"water-delivery/internal/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (s *ProductService) ListActive(ctx context.Context) ([]domain.Product, error) {
	products, err := s.productRepository.ListActive(ctx)
	if err != nil {
		return nil, fmt.Errorf("list active products: %w", err)
	}

	return products, nil
}
