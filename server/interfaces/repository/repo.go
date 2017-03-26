package repository

import mgo "gopkg.in/mgo.v2"

type Query map[string]interface{}
type Change mgo.Change
type Index mgo.Index

type DbHandler interface {
	Insert(name string, obj interface{}) error
	Update(name string, query Query, change Change, result interface{}) error
	UpdateAll(name string, query Query, change Query) (int, error)
	FindOne(name string, query Query, result interface{}) error
	FindAll(name string, query Query, result interface{}, limit int, sort string) error
	Count(name string, query Query) (int, error)
	RemoveOne(name string, query Query) error
	RemoveAll(name string, query Query) error
	Exists(name string, query Query) bool
	DropCollection(name string) error
	DropDatabase() error
	EnsureIndex(name string, index mgo.Index) error
}

type DbRepo struct {
	dbHandlers map[string]DbHandler
	dbHandler  DbHandler
}

type ImageHandler interface {
	Process(image string) (interface{}, error)
}

type MailHandler interface {
	Send(from, to, subject, content string) error
}
