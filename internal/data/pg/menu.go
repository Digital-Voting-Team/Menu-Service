package pg

import (
	"database/sql"
	"gitlab.com/distributed_lab/kit/pgdb"
	"menu-service/internal/data"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
)

const menusTableName = "public.menus"

func NewMenusQ(db *pgdb.DB) data.MenusQ {
	return &menusQ{
		db:        db.Clone(),
		sql:       sq.Select("menus.*").From(menusTableName),
		sqlUpdate: sq.Update(menusTableName).Suffix("returning *"),
	}
}

type menusQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (q *menusQ) New() data.MenusQ {
	return NewMenusQ(q.db)
}

func (q *menusQ) Get() (*data.Menu, error) {
	var result data.Menu
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *menusQ) Select() ([]data.Menu, error) {
	var result []data.Menu
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *menusQ) Update(menu data.Menu) (data.Menu, error) {
	var result data.Menu
	clauses := structs.Map(menu)
	clauses["cafe_id"] = menu.CafeId

	err := q.db.Get(&result, q.sqlUpdate.SetMap(clauses))

	return result, err
}

func (q *menusQ) Transaction(fn func(q data.MenusQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *menusQ) Insert(menu data.Menu) (data.Menu, error) {
	clauses := structs.Map(menu)
	clauses["cafe_id"] = menu.CafeId

	var result data.Menu
	stmt := sq.Insert(menusTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}

func (q *menusQ) Delete(id int64) error {
	stmt := sq.Delete(menusTableName).Where(sq.Eq{"id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *menusQ) Page(pageParams pgdb.OffsetPageParams) data.MenusQ {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}

func (q *menusQ) FilterById(ids ...int64) data.MenusQ {
	q.sql = q.sql.Where(sq.Eq{"id": ids})
	q.sqlUpdate = q.sqlUpdate.Where(sq.Eq{"id": ids})
	return q
}

func (q *menusQ) FilterByCafeId(ids ...int64) data.MenusQ {
	q.sql = q.sql.Where(sq.Eq{"cafe_id": ids})
	return q
}
