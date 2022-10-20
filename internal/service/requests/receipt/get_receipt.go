package requests

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetReceiptRequest struct {
	ReceiptId int64 `url:"-"`
}

func NewGetReceiptRequest(r *http.Request) (GetReceiptRequest, error) {
	request := GetReceiptRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.ReceiptId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
