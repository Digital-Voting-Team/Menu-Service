package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetReceiptListRequest struct {
	pgdb.OffsetPageParams
}

func NewGetReceiptListRequest(r *http.Request) (GetReceiptListRequest, error) {
	var request GetReceiptListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
