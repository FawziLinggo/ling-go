package repository

import (
	"FawziLinggo/ling-go/2.restful-api-go/helper"
	"FawziLinggo/ling-go/2.restful-api-go/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryReposirotyImpl struct {
}

func (repository *CategoryReposirotyImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(name) values(?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)

	// check error
	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	// check error
	helper.PanicIfError(err)

	category.Id = int(id)
	return category

}

func (repository *CategoryReposirotyImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)

	// check error
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryReposirotyImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delet from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)

	// check error
	helper.PanicIfError(err)
}

func (repository *CategoryReposirotyImpl) FindById(ctx context.Context, tx *sql.Tx, CategoryId int) (domain.Category, error) {
	SQL := "select id, name from category where is =?"
	rows, err := tx.QueryContext(ctx, SQL, CategoryId)

	// check error
	helper.PanicIfError(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)

		// Check error
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}
func (repository *CategoryReposirotyImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)

	// Error
	helper.PanicIfError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)

		// Error
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
