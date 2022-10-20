package requests

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetCategoryRequest struct {
	CategoryId int64 `url:"-"`
}

func NewGetCategoryRequest(r *http.Request) (GetCategoryRequest, error) {
	request := GetCategoryRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.CategoryId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
