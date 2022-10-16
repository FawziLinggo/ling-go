package repository

import (
	"FawziLinggo/ling-go/2.restful-api-go/model/domain"
	"context"
	"database/sql"
)

type CategoryReposiroty interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, CategoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
