package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetCategoryListRequest struct {
	pgdb.OffsetPageParams
	FilterCategoryName []string `filter:"category_name"`
	FilterUnit         []string `filter:"unit"`
}

func NewGetCategoryListRequest(r *http.Request) (GetCategoryListRequest, error) {
	var request GetCategoryListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
