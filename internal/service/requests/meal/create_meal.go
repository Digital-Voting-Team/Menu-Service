package requests

import (
	"encoding/json"
	"menu-service/internal/service/helpers"
	"menu-service/resources"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateMealRequest struct {
	Data resources.Meal
}

func NewCreateMealRequest(r *http.Request) (CreateMealRequest, error) {
	var request CreateMealRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateMealRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/attributes/meal_name": validation.Validate(&r.Data.Attributes.MealName, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/price": validation.Validate(&r.Data.Attributes.Price, validation.Required,
			validation.By(helpers.IsFloat)),
		"/data/attributes/amount": validation.Validate(&r.Data.Attributes.Amount, validation.Required,
			validation.By(helpers.IsFloat)),
		"/data/relationships/category/data/id": validation.Validate(&r.Data.Relationships.Category.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
	}).Filter()
}
