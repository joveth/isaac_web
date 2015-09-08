package model

import (
	"gopkg.in/mgo.v2"
)

type Dao struct {
	session *mgo.Session
}

const (
	DbName            = "isaac"
	UserCollection    = "user"
	TopicCollection   = "topic"
	MessageCollection = "gb_messages"
	HistoryCollection = "gb_historys"
	EmailCollection   = "gb_emails"
)

func NewDao() (*Dao, error) {
	session, err := mgo.Dial("mongodb://jov:123456@ds040898.mongolab.com:40898/isaac")
	//session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return &Dao{session}, nil
}
func (d *Dao) Close() {
	d.session.Close()
}
