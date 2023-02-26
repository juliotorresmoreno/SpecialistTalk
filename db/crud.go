package db

import (
	"net/http"
	"sync"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/juliotorresmoreno/SpecialistTalk/helper"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
)

type Crud struct {
	session *model.User
	mu      *sync.Mutex
}

func (that *Crud) getConn() (*Engine, error) {
	conf := configs.GetConfig()
	var conn *Engine
	var err error

	that.mu.Lock()
	if that.session != nil {
		conn, err = GetConnectionPoolWithSession(conf.Database, that.session)
		if err != nil {
			return conn, helper.MakeHTTPError(http.StatusInternalServerError, err)
		}
	} else {
		conn, err = GetConnectionPool()
		if err != nil {
			return conn, helper.MakeHTTPError(http.StatusInternalServerError, err)
		}
	}
	that.mu.Unlock()
	return conn, nil
}

func NewCrud(session *model.User) *Crud {
	return &Crud{session: session, mu: &sync.Mutex{}}
}

func (that *Crud) Get(bean model.Model) error {
	conn, err := that.getConn()
	if err != nil {
		return err
	}
	_, err = conn.Get(bean)
	if err != nil {
		return helper.MakeHTTPError(500, err)
	}
	return nil
}

func (that *Crud) Find(rowSlicePtr interface{}, options model.FindOptions) error {
	conn, err := that.getConn()
	if err != nil {
		return err
	}
	query := conn.Limit(options.Limit, options.Skip)
	if options.OrderBy != nil {
		query = query.OrderBy(options.OrderBy)
	}
	if err = query.Find(rowSlicePtr); err != nil {
		return helper.MakeHTTPError(500, err)
	}
	return nil
}

func (that *Crud) Add(bean model.Model) error {
	if err := bean.Check(); err != nil {
		return helper.HTTPErrorBadRequest
	}
	conn, err := that.getConn()
	if err != nil {
		return err
	}
	_, err = conn.InsertOne(bean)
	if err != nil {
		return helper.MakeHTTPError(500, err)
	}
	return nil
}

func (that *Crud) Update(id interface{}, bean model.Model) error {
	if err := bean.Check(); err != nil {
		return helper.HTTPErrorBadRequest
	}
	conn, err := that.getConn()
	if err != nil {
		return err
	}
	_, err = conn.ID(id).Update(bean)
	if err != nil {
		return helper.MakeHTTPError(500, err)
	}
	return nil
}

func (that *Crud) Delete(id interface{}) error {
	conn, err := that.getConn()
	if err != nil {
		return err
	}
	_, err = conn.Delete(id)
	if err != nil {
		return helper.MakeHTTPError(500, err)
	}
	return nil
}
