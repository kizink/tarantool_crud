package crudapi

import "encoding/json"

type AddRequest struct {
	Key   string          `json:"key" validate:"required"`
	Value json.RawMessage `json:"value" validate:"required"`
}

type UpdateRequest struct {
	Value json.RawMessage `json:"value" validate:"required"`
}

type ItemResponse struct {
	Key   string          `json:"key" validate:"required"`
	Value json.RawMessage `json:"value" validate:"required"`
}
