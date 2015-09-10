package model

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Replay struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	TopicId    string        `bson:"topicid"`
	UName      string        `bson:"uname"`
	Content    string        `bson:"content"`
	CreateDate time.Time     `bson:"createdate"`
}

func (dao *Dao) InserReplay(replay *Replay) error {
	uCollection := dao.session.DB(DBNAME).C(T_REPLAY)
	replay.Id = bson.NewObjectId()
	replay.CreateDate = time.Now()
	_, err := uCollection.Upsert(bson.M{"_id": replay.Id}, replay)
	if err != nil {
		revel.WARN.Printf("Unable to save replay: %v error %v", replay, err)
	}
	return err
}
func (dao *Dao) GetReplays(topicid string) ([]Replay, int) {
	collection := dao.session.DB(DBNAME).C(T_REPLAY)
	replays := []Replay{}
	query := collection.Find(bson.M{"topicid": topicid}).Sort("createdate")
	query.All(&replays)
	return replays, len(replays)
}
