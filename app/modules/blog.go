package modules

import (
	"github.com/revel/revel"
	//	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Blog struct {
	CDate   time.Time
	Title   string
	Content string
	ReadCnt int
	Year    int
}

//新建博客
func (dao *Dao) CreateBlog(blog *Blog) error {
	blogConnection := dao.session.DB(DbName).C(BlogConnection)
	blog.CDate = time.Now()
	blog.Year = blog.CDate.Year()
	blog.ReadCnt = 0
	blog.CDate = time.Now()
	blog.Year = blog.CDate.Year()
	err := blogConnection.Insert(blog) //先根据Id查找，然后更新或插入
	if err != nil {
		revel.WARN.Printf("Unable to save blog:%v error % v", blog, err)
	}
	return err
}

//查询所有博客
func (dao *Dao) FindBlogs() []Blog {
	blogConnection := dao.session.DB(DbName).C(BlogConnection)
	blogs := []Blog{}
	query := blogConnection.Find(bson.M{}).Sort("-cdate").Limit(50) //结果根据cdate倒序
	query.All(&blogs)
	return blogs
}

//前台数据提交校验
func (blog *Blog) Validate(v *revel.Validation) {
	v.Check(blog.Title, revel.Required{}, revel.MinSize{1}, revel.MaxSize{200})
	v.Check(blog.Content, revel.Required{}, revel.MinSize{1})
}
