package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"water-delivery/internal/domain"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) ListActive(ctx context.Context) ([]domain.Product, error) {
	const query = `
		SELECT
			id,
			name,
			description,
			price_cents,
			volume_ml,
			is_active,
			created_at,
			updated_at
		FROM products
		WHERE is_active = TRUE
		ORDER BY id ASC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query active products: %w", err)
	}
	defer rows.Close()

	products := make([]domain.Product, 0)

	for rows.Next() {
		var product domain.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.PriceCents,
			&product.VolumeML,
			&product.IsActive,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan product: %w", err)
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate product rows: %w", err)
	}

	return products, nil
}
