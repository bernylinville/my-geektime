package initialize

import (
	"context"

	"github.com/zkep/mygeektime/internal/global"
	"github.com/zkep/mygeektime/internal/model"
	"github.com/zkep/mygeektime/lib/db"
)

func Gorm(_ context.Context) error {
	g, err := db.NewGORM(
		global.CONF.DB.Driver,
		global.CONF.DB.Source,
		db.MaxIdleConns(global.CONF.DB.MaxIdleConns),
		db.MaxOpenConns(global.CONF.DB.MaxOpenConns),
	)()
	if err != nil {
		return err
	}
	global.DB = g
	if err = g.AutoMigrate(
		&model.User{},
		&model.Task{},
	); err != nil {
		return err
	}
	return nil
}
