package model

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Name       string    `bson:"name"`
	Email      string    `bson:"email"`
	Phone      string    `bson:"phone"`
	Pass       string    `bson:"pass"`
	CreateDate time.Time `bson:"createdate"`
	Logo       string    `bson:"logo"`
	Flag       bool      `bson:"flag"`
}

func (dao *Dao) InserUser(user *User) error {
	uCollection := dao.session.DB(DBNAME).C(T_USER)
	//set the time
	user.CreateDate = time.Now()
	_, err := uCollection.Upsert(bson.M{"email": user.Email}, user)
	if err != nil {
		revel.WARN.Printf("Unable to save user: %v error %v", user, err)
	}
	return err
}
