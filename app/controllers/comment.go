package controllers

import (
	"github.com/revel/revel"
	"myapp/app/modules"
	"strings"
)

type Comment struct {
	App
}

//保存回复
func (c Comment) Save(id string, rcnt int, comment *modules.Comment) revel.Result {
	if len(id) == 0 {
		return c.Redirect(App.Index())
	}
	dao, err := modules.Conn()
	if err != nil { //如果报错
		c.Response.Status = 500
		return c.Redirect(App.Index())
	}
	defer dao.Close()
	blog := dao.GetBlogFromId(id)
	if blog == nil {
		return c.Redirect(App.Index())
	}
	comment.BlogId = blog.Id_
	comment.Content = strings.TrimSpace(comment.Content)
	comment.Email = strings.TrimSpace(comment.Email)
	comment.Validate(c.Validation)
	if c.Validation.HasErrors() {

	}

}
