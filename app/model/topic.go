package model

import (
	"fmt"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"time"
)

type Topic struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	UName      string        `bson:"uname"`
	Title      string        `bson:"title"`
	Body       template.HTML `bson:"body"`
	Tag        string        `bson:"tag"`
	CreateDate time.Time     `bson:"createdate"`
	Status     int           `bson:"status"`
}

func (dao *Dao) InserTopic(topic *Topic) (string, error) {
	collection := dao.session.DB(DbName).C(TopicCollection)
	//set the time
	topic.CreateDate = time.Now()
	topic.Id = bson.NewObjectId()
	topic.Status = 1
	_, err := collection.Upsert(bson.M{"_id": topic.Id}, topic)
	if err != nil {
		revel.WARN.Printf("Unable to save topic: %v error %v", topic, err)
	}
	return topic.Id.Hex(), err
}
func (dao *Dao) FindTopicById(id string) *Topic {
	collection := dao.session.DB(DbName).C(TopicCollection)
	topic := new(Topic)
	query := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)})
	query.One(&topic)
	fmt.Printf("query=%v", query)
	return topic
}
