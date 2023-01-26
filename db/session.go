package db

import (
	"fmt"
	"reflect"

	"github.com/go-xorm/xorm"
	"github.com/juliotorresmoreno/freelive/model"
)

type Session struct {
	permisionQueryRead  string
	permisionQueryWrite string
	user                *model.User
	*xorm.Session
}

func (e *Session) getPermisionQueryRead() string {
	return fmt.Sprintf(e.permisionQueryRead, e.user.Username, e.user.ACL.Group)
}

func (e *Session) getPermisionQueryWrite() string {
	return fmt.Sprintf(e.permisionQueryWrite, e.user.Username, e.user.ACL.Group)
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

var aclType = reflect.ValueOf(model.ACL{}).Type()

func (e *Session) Insert(bean ...interface{}) (int64, error) {
	for _, b := range bean {
		field := reflect.ValueOf(b)
		if field.Kind() != reflect.Ptr {
			field = reflect.ValueOf(&b)
		}
		field = field.Elem().FieldByName("ACL")
		if field.CanSet() && field.Type() == aclType {
			acl, err := model.NewACL(e.user.Username, e.user.ACL.Group)
			if err != nil {
				return 0, err
			}
			field.Set(reflect.ValueOf(acl))
		}
	}

	return e.Session.Insert(bean...)
}

func (e *Session) InsertOne(bean interface{}) (int64, error) {
	field := reflect.ValueOf(bean)
	if field.Kind() != reflect.Ptr {
		field = reflect.ValueOf(&bean)
	}
	field = field.Elem().FieldByName("ACL")
	if field.CanSet() && field.Type() == aclType && field.IsZero() {
		acl, err := model.NewACL(e.user.Username, e.user.ACL.Group)
		if err != nil {
			return 0, err
		}
		field.Set(reflect.ValueOf(acl))
	}
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
