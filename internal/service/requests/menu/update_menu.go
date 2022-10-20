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

type UpdateMenuRequest struct {
	MenuId int64 `url:"-" json:"-"`
	Data   resources.Menu
}

func NewUpdateMenuRequest(r *http.Request) (UpdateMenuRequest, error) {
	request := UpdateMenuRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.MenuId = cast.ToInt64(chi.URLParam(r, "id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *UpdateMenuRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{}).Filter()
}
