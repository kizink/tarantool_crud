package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kizink/tarantool_crud/configs"
	"github.com/kizink/tarantool_crud/internal/crudapi"
	"github.com/kizink/tarantool_crud/pkg/storage"
	"go.uber.org/zap"
)

func main() {
	// config
	conf := configs.LoadConfig()

	// logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Panic(err.Error())
	}
	defer logger.Sync() // flushes buffer, if any
	sugarLog := logger.Sugar()

	// storage
	sugarLog.Info("Initializing storage...")
	db := storage.New(sugarLog, conf)

	// repositories
	sugarLog.Info("Initializing repositories...")
	repo := crudapi.NewTarantoolCrudRepo(db, conf)

	// routers
	r := chi.NewRouter()

	// mount handlers
	crudapi.MountCrudAPIHandlersTo(r, &crudapi.CrudAPIHandlersDeps{
		Repo: repo,
		Log:  sugarLog,
	})

	// init server
	server := http.Server{
		Addr:         ":8081", // добавить в конфиг
		Handler:      r,
		ReadTimeout:  4 * time.Second,  // добавить в конфиг
		WriteTimeout: 4 * time.Second,  // добавить в конфиг
		IdleTimeout:  30 * time.Second, // добавить в конфиг
	}

	// start server
	sugarLog.Info("Starting http server on port 8081...")
	err = server.ListenAndServe()
	if err != nil {
		sugarLog.Panicln("failed to start server: ", err.Error())
	}
}
