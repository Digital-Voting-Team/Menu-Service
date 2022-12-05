package pg

import (
	"database/sql"
	"github.com/Digital-Voting-Team/menu-service/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
)

const categoriesTableName = "public.categories"

func NewCategoriesQ(db *pgdb.DB) data.CategoriesQ {
	return &categoriesQ{
		db:        db.Clone(),
		sql:       sq.Select("categories.*").From(categoriesTableName),
		sqlUpdate: sq.Update(categoriesTableName).Suffix("returning *"),
	}
}

type categoriesQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (q *categoriesQ) New() data.CategoriesQ {
	return NewCategoriesQ(q.db)
}

func (q *categoriesQ) Get() (*data.Category, error) {
	var result data.Category
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *categoriesQ) Select() ([]data.Category, error) {
	var result []data.Category
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *categoriesQ) Update(category data.Category) (data.Category, error) {
	var result data.Category
	clauses := structs.Map(category)
	clauses["category_name"] = category.CategoryName
	clauses["unit"] = category.Unit

	err := q.db.Get(&result, q.sqlUpdate.SetMap(clauses))

	return result, err
}

func (q *categoriesQ) Transaction(fn func(q data.CategoriesQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *categoriesQ) Insert(category data.Category) (data.Category, error) {
	clauses := structs.Map(category)
	clauses["category_name"] = category.CategoryName
	clauses["unit"] = category.Unit

	var result data.Category
	stmt := sq.Insert(categoriesTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}

func (q *categoriesQ) Delete(id int64) error {
	stmt := sq.Delete(categoriesTableName).Where(sq.Eq{"id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *categoriesQ) Page(pageParams pgdb.OffsetPageParams) data.CategoriesQ {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}

func (q *categoriesQ) FilterById(ids ...int64) data.CategoriesQ {
	q.sql = q.sql.Where(sq.Eq{"id": ids})
	q.sqlUpdate = q.sqlUpdate.Where(sq.Eq{"id": ids})
	return q
}

func (q *categoriesQ) FilterByNames(names ...string) data.CategoriesQ {
	q.sql = q.sql.Where(sq.Eq{"category_name": names})
	return q
}

func (q *categoriesQ) FilterByUnits(units ...string) data.CategoriesQ {
	q.sql = q.sql.Where(sq.Eq{"unit": units})
	return q
}
