package modules

import (
	"gopkg.in/mgo.v2"
)

const (
	DbName            = "blog"
	UserCollection    = "user"
	BlogCollection    = "blog"
	CommentCollection = "comment"
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
