package pg

import (
	"Menu-Service/internal/data"
	"database/sql"
	"fmt"
	"gitlab.com/distributed_lab/kit/pgdb"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
)

const mealsTableName = "public.meals"

func NewMealsQ(db *pgdb.DB) data.MealsQ {
	return &mealsQ{
		db:        db.Clone(),
		sql:       sq.Select("meals.*").From(mealsTableName),
		sqlUpdate: sq.Update(mealsTableName).Suffix("returning *"),
	}
}

type mealsQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (q *mealsQ) New() data.MealsQ {
	return NewMealsQ(q.db)
}

func (q *mealsQ) Get() (*data.Meal, error) {
	var result data.Meal
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *mealsQ) Select() ([]data.Meal, error) {
	var result []data.Meal
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *mealsQ) Update(meal data.Meal) (data.Meal, error) {
	var result data.Meal
	clauses := structs.Map(meal)
	clauses["meal_name"] = meal.MealName
	clauses["category_id"] = meal.CategoryId
	clauses["price"] = meal.Price
	clauses["amount"] = meal.Amount

	err := q.db.Get(&result, q.sqlUpdate.SetMap(clauses))

	return result, err
}

func (q *mealsQ) Transaction(fn func(q data.MealsQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *mealsQ) Insert(meal data.Meal) (data.Meal, error) {
	clauses := structs.Map(meal)
	clauses["meal_name"] = meal.MealName
	clauses["category_id"] = meal.CategoryId
	clauses["price"] = meal.Price
	clauses["amount"] = meal.Amount

	var result data.Meal
	stmt := sq.Insert(mealsTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}

func (q *mealsQ) Delete(id int64) error {
	stmt := sq.Delete(mealsTableName).Where(sq.Eq{"id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *mealsQ) Page(pageParams pgdb.OffsetPageParams) data.MealsQ {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}

func (q *mealsQ) FilterById(ids ...int64) data.MealsQ {
	q.sql = q.sql.Where(sq.Eq{"id": ids})
	q.sqlUpdate = q.sqlUpdate.Where(sq.Eq{"id": ids})
	return q
}

func (q *mealsQ) FilterByNames(names ...string) data.MealsQ {
	q.sql = q.sql.Where(sq.Eq{"meal_name": names})
	return q
}

func (q *mealsQ) FilterByPriceFrom(prices ...float64) data.MealsQ {
	stmt := sq.GtOrEq{"price": prices}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *mealsQ) FilterByPriceTo(prices ...float64) data.MealsQ {
	stmt := sq.LtOrEq{"price": prices}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *mealsQ) FilterByAmount(amounts ...float64) data.MealsQ {
	q.sql = q.sql.Where(sq.Eq{"amount": amounts})
	return q
}

func (q *mealsQ) FilterByCategoryId(ids ...float64) data.MealsQ {
	q.sql = q.sql.Where(sq.Eq{"category_id": ids})
	return q
}

func (q *mealsQ) JoinCategory() data.MealsQ {
	stmt := fmt.Sprintf("%s as meals on public.categories.id = meals.category_id",
		mealsTableName)
	q.sql = q.sql.Join(stmt)
	return q
}
