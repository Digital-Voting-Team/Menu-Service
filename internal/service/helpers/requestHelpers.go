package helpers

import (
	"encoding/json"
	"github.com/Digital-Voting-Team/cafe-service/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strconv"
)

func ParseCafeResponse(r *http.Response) (*resources.CafeResponse, error) {
	var response resources.CafeResponse

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return &response, errors.Wrap(err, "failed to unmarshal JwtResponse")
	}

	return &response, nil
}

func ValidateCafe(token, endpoint string, id int64) (*resources.CafeResponse, error) {
	req, err := http.NewRequest("GET", endpoint+strconv.FormatInt(id, 10), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to build new request")
	}
	req.Header.Set("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to send request, endpoint: "+endpoint)
	}

	return ParseCafeResponse(res)
}
