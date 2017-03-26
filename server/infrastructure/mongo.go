package infrastructure

import (
	"time"

	"github.com/jtreu/sleeplog-app/server/interfaces/repository"

	mgo "gopkg.in/mgo.v2"
)

type MongoOptions struct {
	ServerName   string
	DatabaseName string
	DialTimeout  time.Duration
}

type MongoHandler struct {
	db      *mgo.Database
	options *MongoOptions
}

func NewMongoHandler(options *MongoOptions) *MongoHandler {
	handler := &MongoHandler{}
	handler.options = options
	return handler
}

func (handler *MongoHandler) NewSession() {
	mongoOptions := handler.options

	// set default DialTimeout value
	if mongoOptions.DialTimeout <= 0 {
		mongoOptions.DialTimeout = 1 * time.Minute
	}

	session, err := mgo.DialWithTimeout(mongoOptions.ServerName, mongoOptions.DialTimeout)
	if err != nil {
		panic(err)
	}

	handler.db = session.DB(mongoOptions.DatabaseName)
}

func (handler *MongoHandler) FindOne(name string, query repository.Query, result interface{}) error {
	return handler.db.C(name).Find(query).One(result)
}

func (handler *MongoHandler) FindAll(name string, query repository.Query, result interface{}, limit int, sort string) error {
	if sort == "" {
		sort = "-_id"
	}
	return handler.db.C(name).Find(query).Sort(sort).Limit(limit).All(result)
}

func (handler *MongoHandler) Count(name string, query repository.Query) (int, error) {
	return handler.db.C(name).Find(query).Count()
}

func (handler *MongoHandler) Insert(name string, obj interface{}) error {
	return handler.db.C(name).Insert(obj)
}

func (handler *MongoHandler) Update(name string, query repository.Query, change repository.Change, result interface{}) error {
	_, err := handler.db.C(name).Find(query).Apply(mgo.Change(change), result)
	return err
}

func (handler *MongoHandler) UpdateAll(name string, query repository.Query, change repository.Query) (int, error) {
	changeInfo, err := handler.db.C(name).UpdateAll(query, change)
	if changeInfo == nil {
		return 0, err
	}
	return changeInfo.Updated, err
}

func (handler *MongoHandler) RemoveOne(name string, query repository.Query) error {
	return handler.db.C(name).Remove(query)
}

func (handler *MongoHandler) RemoveAll(name string, query repository.Query) error {
	_, err := handler.db.C(name).RemoveAll(query)
	return err
}

func (handler *MongoHandler) DropCollection(name string) error {
	return handler.db.C(name).DropCollection()
}

func (handler *MongoHandler) Exists(name string, query repository.Query) bool {
	var result interface{}
	err := handler.db.C(name).Find(query).One(result)
	return (err == nil)
}

func (handler *MongoHandler) DropDatabase() error {
	return handler.db.DropDatabase()
}

func (handler *MongoHandler) EnsureIndex(name string, index mgo.Index) error {
	return handler.db.C(name).EnsureIndex(index)
}
