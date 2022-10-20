package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteMealRequest struct {
	MealId int64 `url:"-"`
}

func NewDeleteMealRequest(r *http.Request) (DeleteMealRequest, error) {
	request := DeleteMealRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.MealId = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
