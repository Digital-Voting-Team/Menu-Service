package pg

import (
	"database/sql"
	"fmt"
	"gitlab.com/distributed_lab/kit/pgdb"
	"menu-service/internal/data"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
)

const mealMenusTableName = "public.meal_menus"

func NewMealMenusQ(db *pgdb.DB) data.MealMenusQ {
	return &mealMenusQ{
		db:        db.Clone(),
		sql:       sq.Select("meal_menus.*").From(mealMenusTableName),
		sqlUpdate: sq.Update(mealMenusTableName).Suffix("returning *"),
	}
}

type mealMenusQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (q *mealMenusQ) New() data.MealMenusQ {
	return NewMealMenusQ(q.db)
}

func (q *mealMenusQ) Get() (*data.MealMenu, error) {
	var result data.MealMenu
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *mealMenusQ) Select() ([]data.MealMenu, error) {
	var result []data.MealMenu
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *mealMenusQ) Update(mealMenu data.MealMenu) (data.MealMenu, error) {
	var result data.MealMenu
	clauses := structs.Map(mealMenu)
	clauses["meal_id"] = mealMenu.MealId
	clauses["menu_id"] = mealMenu.MenuId

	err := q.db.Get(&result, q.sqlUpdate.SetMap(clauses))

	return result, err
}

func (q *mealMenusQ) Transaction(fn func(q data.MealMenusQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *mealMenusQ) Insert(mealMenu data.MealMenu) (data.MealMenu, error) {
	clauses := structs.Map(mealMenu)
	clauses["meal_id"] = mealMenu.MealId
	clauses["menu_id"] = mealMenu.MenuId

	var result data.MealMenu
	stmt := sq.Insert(mealMenusTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}

func (q *mealMenusQ) Delete(id int64) error {
	stmt := sq.Delete(mealMenusTableName).Where(sq.Eq{"id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *mealMenusQ) Page(pageParams pgdb.OffsetPageParams) data.MealMenusQ {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}

func (q *mealMenusQ) FilterById(ids ...int64) data.MealMenusQ {
	q.sql = q.sql.Where(sq.Eq{"id": ids})
	q.sqlUpdate = q.sqlUpdate.Where(sq.Eq{"id": ids})
	return q
}

func (q *mealMenusQ) FilterByMealId(ids ...int64) data.MealMenusQ {
	q.sql = q.sql.Where(sq.Eq{"meal_id": ids})
	return q
}

func (q *mealMenusQ) FilterByMenuId(ids ...int64) data.MealMenusQ {
	q.sql = q.sql.Where(sq.Eq{"menu_id": ids})
	return q
}

func (q *mealMenusQ) JoinMeal() data.MealMenusQ {
	stmt := fmt.Sprintf("%s as meal_menus on public.meals.id = meal_menus.meal_id",
		mealMenusTableName)
	q.sql = q.sql.Join(stmt)
	return q
}

func (q *mealMenusQ) JoinMenu() data.MealMenusQ {
	stmt := fmt.Sprintf("%s as meal_menus on public.menus.id = meal_menus.menu_id",
		mealMenusTableName)
	q.sql = q.sql.Join(stmt)
	return q
}
