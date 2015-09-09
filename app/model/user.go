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
	Logo       string    `bson:"logo"` //http://7xlhe5.com1.z0.glb.clouddn.com/default.jpg?imageView2/2/w/48/h/48/q/100
	Flag       bool      `bson:"flag"`
	Vip        int       `bson:"vip"`
}

func (dao *Dao) InserUser(user *User) error {
	uCollection := dao.session.DB(DBNAME).C(T_USER)
	//set the time
	user.CreateDate = time.Now()
	_, err := uCollection.Upsert(bson.M{"name": user.Name}, user)
	if err != nil {
		revel.WARN.Printf("Unable to save user: %v error %v", user, err)
	}
	return err
}
func (dao *Dao) UpdateUser(user *User) error {
	uCollection := dao.session.DB(DBNAME).C(T_USER)
	//set the time
	_, err := uCollection.Upsert(bson.M{"name": user.Name}, user)
	if err != nil {
		revel.WARN.Printf("Unable to UpdateUser user: %v error %v", user, err)
	}
	return err
}
func (dao *Dao) GetUserByName(name string) *User {
	collection := dao.session.DB(DBNAME).C(T_USER)
	user := new(User)
	query := collection.Find(bson.M{"name": name})
	query.One(&user)
	return user
}
func (dao *Dao) GetUserByEmail(email string) *User {
	collection := dao.session.DB(DBNAME).C(T_USER)
	user := new(User)
	query := collection.Find(bson.M{"email": email})
	query.One(&user)
	return user
}
func (dao *Dao) GetUserLogoByName(name string) *User {
	collection := dao.session.DB(DBNAME).C(T_USER)
	user := new(User)
	collection.Find(bson.M{"name": name}).Select(bson.M{"logo": 1}).One(&user)
	return user
}
func (dao *Dao) GetUserForLogin(name string, pass string) *User {
	collection := dao.session.DB(DBNAME).C(T_USER)
	user := new(User)
	query := collection.Find(bson.M{"pass": pass, "$or": []bson.M{bson.M{"name": name}, bson.M{"email": name}}})
	query.One(&user)
	return user
}
