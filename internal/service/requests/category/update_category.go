package requests

import (
	"Menu-Service/internal/service/helpers"
	"Menu-Service/resources"
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type UpdateCategoryRequest struct {
	CategoryId int64 `url:"-" json:"-"`
	Data       resources.Category
}

func NewUpdateCategoryRequest(r *http.Request) (UpdateCategoryRequest, error) {
	request := UpdateCategoryRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.CategoryId = cast.ToInt64(chi.URLParam(r, "id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *UpdateCategoryRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/attributes/category_name": validation.Validate(&r.Data.Attributes.CategoryName, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/unit": validation.Validate(&r.Data.Attributes.Unit, validation.Required,
			validation.Length(1, 10)),
	}).Filter()
}
