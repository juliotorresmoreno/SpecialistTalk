package db

import (
	"fmt"
	"math"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"github.com/lib/pq"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"
)

var maxElementSize = int(math.Pow(2, 16))

func NewEngigne(conf *configs.Database) (*xorm.Engine, error) {
	var err error
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

	cacher := caches.NewLRUCacher(caches.NewMemoryStore(), maxElementSize)
	conn.SetDefaultCacher(cacher)

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

	queryRead := fmt.Sprintf("owner = '%v'", user.Username)
	queryWrite := fmt.Sprintf("owner = '%v'", user.Username)

	engine := &Engine{Engine: conn.Engine}
	engine.permisionQueryRead = queryRead
	engine.permisionQueryWrite = queryWrite
	engine.user = user

	return engine, err
}
