package db

import (
	"github.com/go-xorm/xorm"
	"github.com/juliotorresmoreno/freelive/configs"
	"github.com/juliotorresmoreno/freelive/model"
	"github.com/lib/pq"
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

var connectionPool *xorm.Engine

func GetConnectionPool() (*xorm.Engine, error) {
	var err error
	if connectionPool == nil {
		connectionPool, err = NewEngigne(configs.GetConfig().Database)
	}
	return connectionPool, err
}

func GetConnectionPoolWithSession(conf *configs.Database, user *model.User) (*Engine, error) {
	conn, err := GetConnectionPool()

	r := "(acl->>'owner' = '%v' or (acl->'groups'->'%v'->>'read')::boolean is true)"
	w := "(acl->>'owner' = '%v' or (acl->'groups'->'%v'->>'write')::boolean is true)"

	engine := &Engine{Engine: conn}
	engine.permisionQueryRead = r
	engine.permisionQueryWrite = w
	engine.user = user

	return engine, err
}
