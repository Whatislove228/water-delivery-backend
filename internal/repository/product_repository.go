package repository

import (
	"context"

	"water-delivery/internal/domain"
)

type ProductRepository interface {
	ListActive(ctx context.Context) ([]domain.Product, error)
}
