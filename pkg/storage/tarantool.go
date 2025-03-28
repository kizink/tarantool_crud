package storage

import (
	"context"
	"time"

	"github.com/kizink/tarantool_crud/configs"
	"github.com/tarantool/go-tarantool/v2"
	"go.uber.org/zap"
)

type DB struct {
	Conn *tarantool.Connection
}

func New(log *zap.SugaredLogger, conf *configs.Config) *DB {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dialer := tarantool.NetDialer{
		Address: conf.Db.ADDRESS,
		//User:     conf.Db.USER,
		//Password: conf.Db.PASSWORD,
	}

	opts := tarantool.Opts{
		Timeout: time.Second,
	}

	conn, err := tarantool.Connect(ctx, dialer, opts)
	if err != nil {
		log.Panic("Connection refused:" + err.Error())
	}

	return &DB{Conn: conn}
}
