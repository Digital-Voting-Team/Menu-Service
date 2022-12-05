package requests

import (
	"encoding/json"
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	"github.com/Digital-Voting-Team/menu-service/resources"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateMealMenuRequest struct {
	Data resources.MealMenu
}

func NewCreateMealMenuRequest(r *http.Request) (CreateMealMenuRequest, error) {
	var request CreateMealMenuRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateMealMenuRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/relationships/meal/data/id": validation.Validate(&r.Data.Relationships.Meal.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
		"/data/relationships/menu/data/id": validation.Validate(&r.Data.Relationships.Menu.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
	}).Filter()
}
