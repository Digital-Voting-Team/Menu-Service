package pg

import (
	"Menu-Service/internal/data"
	"database/sql"
	"fmt"
	"gitlab.com/distributed_lab/kit/pgdb"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
)

const receiptsTableName = "public.receipts"

func NewReceiptsQ(db *pgdb.DB) data.ReceiptsQ {
	return &receiptsQ{
		db:        db.Clone(),
		sql:       sq.Select("receipts.*").From(receiptsTableName),
		sqlUpdate: sq.Update(receiptsTableName).Suffix("returning *"),
	}
}

type receiptsQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (q *receiptsQ) New() data.ReceiptsQ {
	return NewReceiptsQ(q.db)
}

func (q *receiptsQ) Get() (*data.Receipt, error) {
	var result data.Receipt
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *receiptsQ) Select() ([]data.Receipt, error) {
	var result []data.Receipt
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *receiptsQ) Update(receipt data.Receipt) (data.Receipt, error) {
	var result data.Receipt
	clauses := structs.Map(receipt)
	clauses["meal_id"] = receipt.MealId
	clauses["ingredient_id"] = receipt.IngredientId
	clauses["quantity"] = receipt.Quantity

	err := q.db.Get(&result, q.sqlUpdate.SetMap(clauses))

	return result, err
}

func (q *receiptsQ) Transaction(fn func(q data.ReceiptsQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *receiptsQ) Insert(receipt data.Receipt) (data.Receipt, error) {
	clauses := structs.Map(receipt)
	clauses["meal_id"] = receipt.MealId
	clauses["ingredient_id"] = receipt.IngredientId
	clauses["quantity"] = receipt.Quantity

	var result data.Receipt
	stmt := sq.Insert(receiptsTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}

func (q *receiptsQ) Delete(id int64) error {
	stmt := sq.Delete(receiptsTableName).Where(sq.Eq{"id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *receiptsQ) Page(pageParams pgdb.OffsetPageParams) data.ReceiptsQ {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}

func (q *receiptsQ) FilterById(ids ...int64) data.ReceiptsQ {
	q.sql = q.sql.Where(sq.Eq{"id": ids})
	q.sqlUpdate = q.sqlUpdate.Where(sq.Eq{"id": ids})
	return q
}

func (q *receiptsQ) FilterByMealId(ids ...int64) data.ReceiptsQ {
	q.sql = q.sql.Where(sq.Eq{"meal_id": ids})
	return q
}

func (q *receiptsQ) FilterByIngredientId(ids ...int64) data.ReceiptsQ {
	q.sql = q.sql.Where(sq.Eq{"ingredient_id": ids})
	return q
}

func (q *receiptsQ) FilterByQuantityFrom(quantities ...int64) data.ReceiptsQ {
	stmt := sq.GtOrEq{"quantity": quantities}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *receiptsQ) FilterByQuantityTo(quantities ...int64) data.ReceiptsQ {
	stmt := sq.LtOrEq{"quantity": quantities}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *receiptsQ) JoinMeal() data.ReceiptsQ {
	stmt := fmt.Sprintf("%s as receipts on public.meals.id = receipts.meal_id",
		receiptsTableName)
	q.sql = q.sql.Join(stmt)
	return q
}
