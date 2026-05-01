package domain

import "time"

type Product struct {
	ID          int64
	Name        string
	Description *string
	PriceCents  int
	VolumeML    int
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
