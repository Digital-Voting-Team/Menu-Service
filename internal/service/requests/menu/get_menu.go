package requests

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetMenuRequest struct {
	MenuId int64 `url:"-"`
}

func NewGetMenuRequest(r *http.Request) (GetMenuRequest, error) {
	request := GetMenuRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.MenuId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
