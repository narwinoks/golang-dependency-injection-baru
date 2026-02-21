package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

//	func NewCategoryRepository() CategoryRepository {
//		return &CategoryRepositoryImpl{}
//	}
func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "insert into category(name) value(?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(id)
	return category
}
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err)
}
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sql := "select id,name from category where id = ?"
	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)
	category := domain.Category{}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		fmt.Println("DEBUG ERROR FALSE")
		return category, nil
	} else {
		fmt.Println("DEBUG ERROR TRUE")
		return category, errors.New("category is not found")
	}
}
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id,name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
