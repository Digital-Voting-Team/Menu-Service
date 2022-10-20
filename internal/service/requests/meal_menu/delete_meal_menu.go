package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteMealMenuRequest struct {
	MealMenuId int64 `url:"-"`
}

func NewDeleteMealMenuRequest(r *http.Request) (DeleteMealMenuRequest, error) {
	request := DeleteMealMenuRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.MealMenuId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
