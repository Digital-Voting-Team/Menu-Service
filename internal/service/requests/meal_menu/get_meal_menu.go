package requests

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetMealMenuRequest struct {
	MealMenuId int64 `url:"-"`
}

func NewGetMealMenuRequest(r *http.Request) (GetMealMenuRequest, error) {
	request := GetMealMenuRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.MealMenuId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
