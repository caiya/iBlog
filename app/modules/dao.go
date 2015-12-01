package modules

import (
	//	"fmt"
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
)

const (
	DbName         = "blog"
	UserConnection = "user"
	BlogConnection = "blog"
)

type Dao struct {
	session *mgo.Session
}

func Conn() (*Dao, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return &Dao{session}, nil
}

func (dao *Dao) Close() {
	dao.session.Close()
}
