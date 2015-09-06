package db
import (
	 "gopkg.in/mgo.v2"
	 "time"
)
type User struct {
        Name string
        Email string
        Phone string
        Pass string
        CreateDate time.Time
        Logo string
}

type Dao struct {
	session *mgo.Session
}
func NewDao() (*Dao, error){
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
