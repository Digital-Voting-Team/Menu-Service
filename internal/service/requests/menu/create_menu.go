package requests

import (
	"Menu-Service/internal/service/helpers"
	"Menu-Service/resources"
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateMenuRequest struct {
	Data resources.Menu
}

func NewCreateMenuRequest(r *http.Request) (CreateMenuRequest, error) {
	var request CreateMenuRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateMenuRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{}).Filter()
}
