package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteReceiptRequest struct {
	ReceiptId int64 `url:"-"`
}

func NewDeleteReceiptRequest(r *http.Request) (DeleteReceiptRequest, error) {
	request := DeleteReceiptRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.ReceiptId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
