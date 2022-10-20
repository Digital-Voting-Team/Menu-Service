package requests

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetMealRequest struct {
	MealId int64 `url:"-"`
}

func NewGetMealRequest(r *http.Request) (GetMealRequest, error) {
	request := GetMealRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.MealId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
