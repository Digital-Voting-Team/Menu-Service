package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetMealListRequest struct {
	pgdb.OffsetPageParams
	FilterMealName []string `filter:"meal_name"`
}

func NewGetMealListRequest(r *http.Request) (GetMealListRequest, error) {
	var request GetMealListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
