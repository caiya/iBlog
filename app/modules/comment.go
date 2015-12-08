package modules

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Comment struct {
	Id_     string `bson:"_id"`
	BlogId  string `bson:"_id"`
	Email   string
	CDate   time.Time
	Content string
}

func (comment *Comment) Validate(v *revel.Validation) {
	v.Check(comment.Email, revel.Required{}, revel.MaxSize{50})
	v.Check(comment.Content, revel.Required{}, revel.MinSize{1}, revel.MaxSize{200})
}

func (dao *Dao) InsertComment(comment *Comment) error {
	commentCollection := dao.session.DB(DbName).C(CommentCollection)
	comment.CDate = time.Now()
	err := commentCollection.Insert(comment)
	if err != nil {
		revel.WARN.Printf("Unable to save Comment: %v error %v", comment, err)
	}
	return err
}

func (dao *Dao) getCommentsFromBlogId(id string) []Comment {
	commentCollection := dao.session.DB(DbName).C(CommentCollection)
	comments := []Comment{}
	query := commentCollection.Find(bson.M{"BlogId", id}).Sort("CDate")
	query.All(comments)
	return comments
}
