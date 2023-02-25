package db

import (
	"reflect"

	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"xorm.io/xorm"
)

type Session struct {
	permisionQueryRead  string
	permisionQueryWrite string
	limit               int
	skip                int
	user                *model.User
	*xorm.Session
}

func (e *Session) getPermisionQueryRead() string {
	return e.permisionQueryRead
}

func (e *Session) getPermisionQueryWrite() string {
	return e.permisionQueryWrite
}

func (e *Session) Get(bean interface{}) (bool, error) {
	if e.permisionQueryRead != "" {
		return e.Session.Where(e.getPermisionQueryRead()).Get(bean)
	}
	return e.Session.Get(bean)
}

func (e *Session) Update(bean interface{}, condiBean ...interface{}) (int64, error) {
	if e.permisionQueryWrite != "" {
		return e.Session.Where(e.getPermisionQueryWrite()).Update(bean, condiBean...)
	}
	return e.Session.Update(bean, condiBean...)
}

func (e *Session) SessionWithACL() *xorm.Session {
	if e.permisionQueryWrite != "" {
		return e.Session.Where(e.getPermisionQueryWrite())
	}
	return e.Session
}

func (e *Session) setOwner(field reflect.Value) {
	field = field.Elem().FieldByName("Owner")
	value := field.String()
	if field.CanSet() && e.user != nil && value == "" {
		value := reflect.ValueOf(e.user.Username)
		field.Set(value)
	}
}

func (e *Session) Insert(bean ...interface{}) (int64, error) {
	for _, b := range bean {
		field := reflect.ValueOf(b)
		if field.Kind() != reflect.Ptr {
			field = reflect.ValueOf(&b)
		}
		e.setOwner(field)
	}

	return e.Session.Insert(bean...)
}

func (e *Session) InsertOne(bean interface{}) (int64, error) {
	field := reflect.ValueOf(bean)
	if field.Kind() != reflect.Ptr {
		field = reflect.ValueOf(&bean)
	}
	e.setOwner(field)
	return e.Session.InsertOne(bean)
}

func (e *Session) Delete(bean interface{}) (int64, error) {
	if e.permisionQueryWrite != "" {
		return e.Session.Where(e.getPermisionQueryWrite()).Delete(bean)
	}
	return e.Session.Update(bean)
}

func (e *Session) Find(rowsSlicePtr interface{}, condiBean ...interface{}) error {
	if e.permisionQueryRead != "" {
		return e.Session.Where(e.getPermisionQueryRead()).Find(rowsSlicePtr, condiBean...)
	}
	return e.Session.Find(rowsSlicePtr, condiBean...)
}

func (e Session) Count(bean ...interface{}) (int64, error) {
	if e.permisionQueryRead != "" {
		return e.Session.Where(e.getPermisionQueryRead()).Count(bean...)
	}
	return e.Session.Count(bean...)
}

func (e Session) Begin() error {
	return e.Session.Begin()
}

func (e Session) Commit() error {
	return e.Session.Commit()
}

func (e Session) Select(str string) *Session {
	e.Session = e.Session.Select(str)
	return &e
}

func (e Session) Where(query interface{}, args ...interface{}) *Session {
	e.Session = e.Session.Where(query, args...)
	if e.limit > 0 {
		e.Limit(e.limit, e.skip)
	}
	return &e
}

func (e Session) Or(query interface{}, args ...interface{}) *Session {
	e.Session = e.Session.Or(query, args...)
	return &e
}

func (e Session) And(query interface{}, args ...interface{}) *Session {
	e.Session = e.Session.And(query, args...)
	return &e
}

func (e Session) Table(tableNameOrBean interface{}) *Session {
	e.Session = e.Session.Table(tableNameOrBean)
	return &e
}
