package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteMenuRequest struct {
	MenuId int64 `url:"-"`
}

func NewDeleteMenuRequest(r *http.Request) (DeleteMenuRequest, error) {
	request := DeleteMenuRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.MenuId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
