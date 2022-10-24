package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetMealMenuListRequest struct {
	pgdb.OffsetPageParams
	FilterMealId []int64 `filter:"meal_id"`
	FilterMenuId []int64 `filter:"menu_id"`
}

func NewGetMealMenuListRequest(r *http.Request) (GetMealMenuListRequest, error) {
	var request GetMealMenuListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
