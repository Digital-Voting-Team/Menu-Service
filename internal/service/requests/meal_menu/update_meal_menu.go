package requests

import (
	"encoding/json"
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	"github.com/Digital-Voting-Team/menu-service/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type UpdateMealMenuRequest struct {
	MealMenuId int64 `url:"-" json:"-"`
	Data       resources.MealMenu
}

func NewUpdateMealMenuRequest(r *http.Request) (UpdateMealMenuRequest, error) {
	request := UpdateMealMenuRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.MealMenuId = cast.ToInt64(chi.URLParam(r, "id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *UpdateMealMenuRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/relationships/meal/data/id": validation.Validate(&r.Data.Relationships.Meal.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
		"/data/relationships/menu/data/id": validation.Validate(&r.Data.Relationships.Menu.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
	}).Filter()
}
