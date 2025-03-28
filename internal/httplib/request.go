package httplib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var ErrIncorrectBody = fmt.Errorf("request body is incorrect")

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	var payload T
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		JsonResponse(w,
			ErrorResponse{Err: ErrIncorrectBody.Error()},
			http.StatusBadRequest)
		return nil, fmt.Errorf("%s", "HandleBody:Decode err: "+err.Error())
	}

	validate := validator.New()
	err = validate.Struct(payload)
	if err != nil {
		JsonResponse(w,
			ErrorResponse{Err: ErrIncorrectBody.Error()},
			http.StatusBadRequest)
		return nil, fmt.Errorf("%s", "HandleBody:validator err: "+err.Error())
	}
	return &payload, nil
}
