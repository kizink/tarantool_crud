package crudapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kizink/tarantool_crud/internal/httplib"
	"go.uber.org/zap"
)

type CrudAPIHandlersDeps struct {
	Repo Repo
	Log  *zap.SugaredLogger
}

type crudAPIHandlers struct {
	repo Repo
	log  *zap.SugaredLogger
}

func MountCrudAPIHandlersTo(router *chi.Mux, deps *CrudAPIHandlersDeps) {
	handler := &crudAPIHandlers{
		repo: deps.Repo,
		log:  deps.Log,
	}

	router.Post("/kv", handler.Add)
	router.Get("/kv/{id}", handler.Get)
	router.Put("/kv/{id}", handler.Update)
	router.Delete("/kv/{id}", handler.Delete)
}

// POST /kv body: {key: "test", "value": {SOME ARBITRARY JSON}}
func (h *crudAPIHandlers) Add(w http.ResponseWriter, r *http.Request) {
	req, err := httplib.HandleBody[AddRequest](w, r)
	if err != nil {
		h.log.Info(err.Error())
		return
	}

	item := &Item{
		Key:   req.Key,
		Value: req.Value,
	}

	item, err = h.repo.Add(item)
	if err != nil {
		h.log.Info(err.Error())
		if err == ErrAlreadyExistsKey {
			httplib.JsonResponse(w,
				httplib.ErrorResponse{
					Err: err.Error(),
				},
				http.StatusConflict)
		} else {
			httplib.JsonResponse(w, "",
				http.StatusInternalServerError)
		}
		return
	}

	httplib.JsonResponse(w,
		ItemResponse{Key: item.Key, Value: item.Value},
		http.StatusCreated)
}

// GET kv/{id}
func (h *crudAPIHandlers) Get(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "id")
	item, err := h.repo.GetByKey(key)
	if err != nil {
		h.log.Info(err.Error())

		if err == ErrNoTupleWithThisKey {
			httplib.JsonResponse(w,
				httplib.ErrorResponse{
					Err: err.Error(),
				},
				http.StatusNotFound)
		} else {
			httplib.JsonResponse(w, "",
				http.StatusInternalServerError)
		}
		return
	}

	httplib.JsonResponse(w,
		ItemResponse{Key: item.Key, Value: item.Value},
		http.StatusOK)
}

// PUT kv/{id} body: {"value": {SOME ARBITRARY JSON}}
func (h *crudAPIHandlers) Update(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "id")

	req, err := httplib.HandleBody[UpdateRequest](w, r)
	if err != nil {
		h.log.Info(err.Error())
		return
	}

	item := &Item{
		Key:   key,
		Value: req.Value,
	}

	item, err = h.repo.Update(item)
	if err != nil {
		h.log.Info(err.Error())

		if err == ErrNoTupleWithThisKey {
			httplib.JsonResponse(w,
				httplib.ErrorResponse{
					Err: err.Error(),
				},
				http.StatusNotFound)
		} else {
			httplib.JsonResponse(w, "",
				http.StatusInternalServerError)
		}
		return
	}

	httplib.JsonResponse(w,
		ItemResponse{Key: item.Key, Value: item.Value},
		http.StatusOK)
}

// DELETE kv/{id}
func (h *crudAPIHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "id")

	err := h.repo.Delete(key)
	if err != nil {
		h.log.Info(err.Error())

		if err == ErrNoTupleWithThisKey {
			httplib.JsonResponse(w,
				httplib.ErrorResponse{
					Err: err.Error(),
				},
				http.StatusNotFound)
		} else {
			httplib.JsonResponse(w, "",
				http.StatusInternalServerError)
		}
		return
	}
}
