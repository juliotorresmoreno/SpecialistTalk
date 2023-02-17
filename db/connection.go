package db

import (
	"fmt"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"github.com/lib/pq"
	"xorm.io/xorm"
)

func NewEngigne(conf *configs.Database) (*xorm.Engine, error) {
	dsn := conf.DSN
	if conf.Driver == "postgres" {
		dsn, _ = pq.ParseURL(conf.DSN)
	}
	conn, err := xorm.NewEngine(conf.Driver, dsn)
	if err != nil {
		return conn, err
	}
	conn.SetMaxOpenConns(conf.MaxOpenConns)
	conn.SetMaxIdleConns(conf.MaxIdleConns)
	return conn, err
}

var connectionPool *Engine

func GetConnectionPool() (*Engine, error) {
	var err error
	if connectionPool == nil {
		conn, err := NewEngigne(configs.GetConfig().Database)
		if err != nil {
			return connectionPool, err
		}
		connectionPool = &Engine{Engine: conn}
	}
	return connectionPool, err
}

func GetConnectionPoolWithSession(conf *configs.Database, user *model.User) (*Engine, error) {
	conn, err := GetConnectionPool()

	r := fmt.Sprintf("owner = '%v'", user.Username)
	w := fmt.Sprintf("owner = '%v'", user.Username)

	engine := &Engine{Engine: conn.Engine}
	engine.permisionQueryRead = r
	engine.permisionQueryWrite = w
	engine.user = user

	return engine, err
}
