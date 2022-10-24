package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetMenuListRequest struct {
	pgdb.OffsetPageParams
	FilterCafeId []int64 `filter:"cafe_id"`
}

func NewGetMenuListRequest(r *http.Request) (GetMenuListRequest, error) {
	var request GetMenuListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
