package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetReceiptListRequest struct {
	pgdb.OffsetPageParams
	FilterMealId       []int64 `filter:"meal_id"`
	FilterIngredientId []int64 `filter:"ingredient_id"`
	FilterQuantityFrom []int64 `filter:"quantity_from"`
	FilterQuantityTo   []int64 `filter:"quantity_to"`
}

func NewGetReceiptListRequest(r *http.Request) (GetReceiptListRequest, error) {
	var request GetReceiptListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
