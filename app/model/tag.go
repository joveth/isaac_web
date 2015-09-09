package model

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type Tag struct {
	Id   int    `bson:"_id,omitempty"`
	Name string `bson:"uname"`
}

func (dao *Dao) InserTag(tag *Tag) error {
	uCollection := dao.session.DB(DBNAME).C(T_TAG)
	_, err := uCollection.Upsert(bson.M{"_id": tag.Id}, tag)
	if err != nil {
		revel.WARN.Printf("Unable to save tag: %v error %v", tag, err)
	}
	return err
}
func (dao *Dao) GetTag(id int) *Tag {
	collection := dao.session.DB(DBNAME).C(T_TAG)
	tag := new(Tag)
	query := collection.Find(bson.M{"_id": id})
	query.One(&tag)
	revel.INFO.Printf("The GetTag(%d) result: %v", id, tag)
	return tag
}

func (dao *Dao) GetTags() []Tag {
	collection := dao.session.DB(DBNAME).C(T_TAG)
	tags := []Tag{}
	query := collection.Find(bson.M{})
	query.All(&tags)
	return tags
}
