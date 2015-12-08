package modules

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Blog struct {
	Id_     string `bson:"_id"`
	CDate   time.Time
	Title   string
	Content string
	ReadCnt int
	Year    int
	Author  string
}

//新建博客
func (dao *Dao) CreateBlog(blog *Blog) error {
	BlogCollection := dao.session.DB(DbName).C(BlogCollection)
	blog.CDate = time.Now()
	blog.Year = blog.CDate.Year()
	blog.ReadCnt = 0
	blog.CDate = time.Now()
	blog.Id_ = bson.NewObjectId().Hex()
	blog.Author = "匿名"
	blog.Year = blog.CDate.Year()
	err := BlogCollection.Insert(blog) //先根据Id查找，然后更新或插入
	if err != nil {
		revel.WARN.Printf("Unable to save blog:%v error % v", blog, err)
	}
	return err
}

//获取Title
func (blog *Blog) GetTitle() string {
	if len(blog.Title) > 100 {
		return blog.Title[:100]
	}
	return blog.Title
}

//获取Content
func (blog *Blog) GetContent() string {
	if len(blog.Content) > 500 {
		return blog.Content[:500]
	}
	return blog.Content
}

//查询所有博客
func (dao *Dao) FindBlogs() []Blog {
	BlogCollection := dao.session.DB(DbName).C(BlogCollection)
	blogs := []Blog{}
	query := BlogCollection.Find(bson.M{}).Sort("-cdate").Limit(50) //结果根据cdate倒序
	query.All(&blogs)
	return blogs
}

//前台数据提交校验
func (blog *Blog) Validate(v *revel.Validation) {
	v.Check(blog.Title, revel.Required{}, revel.MinSize{1}, revel.MaxSize{200})
	v.Check(blog.Content, revel.Required{}, revel.MinSize{1})
}

//根据id查询Blog对象
func (dao *Dao) GetBlogFromId(id string) *Blog {
	BlogCollection := dao.session.DB(DbName).C(BlogCollection)
	blog := new(Blog)
	query := BlogCollection.Find(bson.M{"_id": id})
	query.One(blog)
	return blog
}

func (dao *Dao) UpdateBlogById(id string, blog *Blog) {
	blogCollection := dao.session.DB(DbName).C(BlogCollection)
	err := blogCollection.Update(bson.M{"_id": id}, blog)
	if err != nil {
		revel.WARN.Printf("Unable to update blog: %v error %v", blog, err)
	}
}
