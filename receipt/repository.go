package receipt

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	queryCreateTable = `CREATE TABLE IF NOT EXISTS public.receipt
	(
		id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
		meal integer NOT NULL,
		ingredient integer NOT NULL,
		quantity integer NOT NULL,
		CONSTRAINT receipt_pkey PRIMARY KEY (id),
		CONSTRAINT meal FOREIGN KEY (meal)
			REFERENCES public.meals (id) MATCH SIMPLE
			ON UPDATE NO ACTION
			ON DELETE NO ACTION
	)
	
	TABLESPACE pg_default;
	
	ALTER TABLE IF EXISTS public.receipt
    OWNER to postgres;`

	queryDeleteTable = `DROP TABLE public.receipt`

	queryInsert = `INSERT INTO public.receipt(
	meal, ingredient, quantity)
	VALUES ($1, $2, $3) RETURNING id;`

	querySelect = `SELECT * FROM public.receipt;`

	queryUpdate = `UPDATE public.receipt
	SET meal=$2, ingredient=$3, quantity=$4
	WHERE id=$1;`

	queryDelete = `DELETE FROM public.receipt
	WHERE id=$1;`

	queryCleanDb = `DELETE FROM public.receipt;`

	queryResetCounter = `alter sequence receipt_id_seq restart with 1`
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Insert(receipt *Receipt) (int, error) {
	rows, err := repo.db.Queryx(queryInsert, receipt.MealId, receipt.IngredientId, receipt.Quantity)
	defer rows.Close()
	id := -1
	if err != nil {
		return id, err
	}

	rows.Next()
	err = rows.Scan(&id)
	return id, nil
}

func (repo *Repository) CreateTable() error {
	_, err := repo.db.Exec(queryCreateTable)
	return err
}

func (repo *Repository) DeleteTable() error {
	_, err := repo.db.Exec(queryDeleteTable)
	return err
}

func (repo *Repository) Select() ([]Receipt, error) {
	rows, err := repo.db.Queryx(querySelect)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	receipt := Receipt{}
	var receiptArray []Receipt
	for rows.Next() {
		err = rows.StructScan(&receipt)
		if err != nil {
			return nil, err
		}
		receiptArray = append(receiptArray, receipt)
	}
	return receiptArray, err
}

func (repo *Repository) Delete(id int) error {
	_, err := repo.db.Exec(queryDelete, id)
	return err
}

func (repo *Repository) Update(id int, receipt *Receipt) error {
	_, err := repo.db.Queryx(queryUpdate, id, receipt.MealId, receipt.IngredientId, receipt.Quantity)
	return err
}

func (repo *Repository) Clean() error {
	_, err := repo.db.Exec(queryCleanDb)
	return err
}

func (repo *Repository) ResetCounter() error {
	_, err := repo.db.Exec(queryResetCounter)
	return err
}
