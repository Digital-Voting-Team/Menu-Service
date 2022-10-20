package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteCategoryRequest struct {
	CategoryId int64 `url:"-"`
}

func NewDeleteCategoryRequest(r *http.Request) (DeleteCategoryRequest, error) {
	request := DeleteCategoryRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.CategoryId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
