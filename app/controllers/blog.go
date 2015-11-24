package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"myapp/app/modules"
	"strings"
)

type WBlog struct {
	App
}

func (c WBlog) Save(blog *modules.Blog) revel.Result {
	blog.Title = strings.TrimSpace(blog.Title)
	blog.Content = strings.TrimSpace(blog.Content)
	blog.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		fmt.Println(c.Validation)
		return c.Redirect(App.Blog)
	}
	dao, err := modules.Conn()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	err = dao.CreateBlog(blog)
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	return c.Redirect(App.Index)
}
