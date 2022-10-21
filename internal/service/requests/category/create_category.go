package requests

import (
	"Menu-Service/internal/service/helpers"
	"Menu-Service/resources"
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateCategoryRequest struct {
	Data resources.Category
}

func NewCreateCategoryRequest(r *http.Request) (CreateCategoryRequest, error) {
	var request CreateCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateCategoryRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/attributes/category_name": validation.Validate(&r.Data.Attributes.CategoryName, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/unit": validation.Validate(&r.Data.Attributes.Unit, validation.Required,
			validation.Length(1, 10)),
	}).Filter()
}
