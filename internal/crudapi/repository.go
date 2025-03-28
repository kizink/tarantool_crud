package crudapi

import (
	"encoding/json"
	"fmt"
)

var (
	ErrAlreadyExistsKey            = fmt.Errorf("tuple with this key already exists")
	ErrNoTupleWithThisKey          = fmt.Errorf("there isn't tuple with this key")
	ErrUnexpectedNumberOfTupleElem = fmt.Errorf("unexpected number of elements in tuple")
)

type Item struct {
	Key   string
	Value json.RawMessage
}

type Repo interface {
	Add(item *Item) (*Item, error)
	Update(item *Item) (*Item, error)
	GetByKey(key string) (*Item, error)
	Delete(key string) error
}
