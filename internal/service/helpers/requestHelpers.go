package helpers

import (
	"encoding/json"
	cafeResources "github.com/Digital-Voting-Team/cafe-service/resources"
	warehouseResources "github.com/Digital-Voting-Team/warehouse-service/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strconv"
)

func ParseCafeResponse(r *http.Response) (*cafeResources.CafeResponse, error) {
	var response cafeResources.CafeResponse

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return &response, errors.Wrap(err, "failed to unmarshal CafeResponse")
	}

	return &response, nil
}

func ParseIngredientResponse(r *http.Response) (*warehouseResources.IngredientResponse, error) {
	var response warehouseResources.IngredientResponse

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return &response, errors.Wrap(err, "failed to unmarshal IngredientResponse")
	}

	return &response, nil
}

func ValidateCafe(token, endpoint string, id int64) (*cafeResources.CafeResponse, error) {
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

func ValidateIngredient(token, endpoint string, id int64) (*warehouseResources.IngredientResponse, error) {
	req, err := http.NewRequest("GET", endpoint+strconv.FormatInt(id, 10), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to build new request")
	}
	req.Header.Set("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to send request, endpoint: "+endpoint)
	}

	return ParseIngredientResponse(res)
}
