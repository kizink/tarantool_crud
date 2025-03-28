package crudapi

import (
	"fmt"
	"strings"

	"github.com/kizink/tarantool_crud/configs"
	"github.com/kizink/tarantool_crud/pkg/storage"
	"github.com/tarantool/go-tarantool/v2"
)

type TarantoolCrudRepo struct {
	Db         *storage.DB
	SPASE_NAME string
}

func NewTarantoolCrudRepo(db *storage.DB, conf *configs.Config) Repo {
	return &TarantoolCrudRepo{
		Db:         db,
		SPASE_NAME: conf.Db.SPACE_NAME,
	}
}

func (repo *TarantoolCrudRepo) Add(item *Item) (*Item, error) {
	req := tarantool.NewInsertRequest(repo.SPASE_NAME).
		Tuple([]interface{}{item.Key, string(item.Value)})
	future := repo.Db.Conn.Do(req)

	itemPtr, err := getItemFrom(future)
	if err != nil {
		// yeah, this is a bad way, but I didn't have time
		// to dive deep into the documentation
		if strings.Contains(err.Error(), "Duplicate key exists in unique index") {
			return nil, ErrAlreadyExistsKey
		}
		return nil, err
	}
	return itemPtr, nil
}

func (repo *TarantoolCrudRepo) Update(item *Item) (*Item, error) {
	req := tarantool.NewUpdateRequest(repo.SPASE_NAME).
		Key(tarantool.StringKey{S: item.Key}).
		Operations(tarantool.NewOperations().Assign(1, string(item.Value)))
	future := repo.Db.Conn.Do(req)

	itemPtr, err := getItemFrom(future)
	if err != nil {
		return nil, err
	}
	return itemPtr, nil
}

func (repo *TarantoolCrudRepo) GetByKey(key string) (*Item, error) {
	req := tarantool.NewSelectRequest(repo.SPASE_NAME).
		Limit(1).
		Iterator(tarantool.IterEq).
		Key([]interface{}{key})

	future := repo.Db.Conn.Do(req)
	itemPtr, err := getItemFrom(future)
	if err != nil {
		return nil, err
	}
	return itemPtr, nil
}

func (repo *TarantoolCrudRepo) Delete(key string) error {
	req := tarantool.NewDeleteRequest(repo.SPASE_NAME).
		Index("primary").
		Key([]interface{}{key})

	future := repo.Db.Conn.Do(req)

	_, err := getItemFrom(future)
	if err != nil {
		return err
	}
	return nil
}

func getItemFrom(f *tarantool.Future) (*Item, error) {
	var items []Item

	err := f.GetTyped(&items)
	if err != nil {
		return nil, fmt.Errorf("%s",
			"getItemFrom: f.GetTyped(&items): "+err.Error())
	}

	if len(items) == 0 {
		return nil, ErrNoTupleWithThisKey
	}

	if len(items) == 1 {
		return &Item{
			Key:   items[0].Key,
			Value: items[0].Value,
		}, nil
	}

	return nil, ErrUnexpectedNumberOfTupleElem
}
