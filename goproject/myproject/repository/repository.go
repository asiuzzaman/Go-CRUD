package repsitory

import (
	"context"

	"goproject/myproject/models"
)

// PostRepo explain...
type PostRepo interface {
	Fetch(ctx context.Context, num int64) ([]*models.Pack, error)
	GetByID(ctx context.Context, id int64) (*models.Pack, error)
	Create(ctx context.Context, p *models.Pack) (int64, error)
	Update(ctx context.Context, p *models.Pack) (*models.Pack, error)
	Delete(ctx context.Context, id int64) (bool, error)
}