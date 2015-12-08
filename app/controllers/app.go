package controllers

import (
	"github.com/revel/revel"
	"myapp/app/modules"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	dao, err := modules.Conn()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	blogs := dao.FindBlogs()
	return c.Render(blogs)
}

func (c App) Blog() revel.Result {
	return c.Render()
}

func (c App) BlogInfo(id string, rcnt int) revel.Result {
	dao, err := modules.Conn()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	blog := dao.GetBlogFromId(id)
	if blog.ReadCnt == rcnt {
		blog.ReadCnt = rcnt + 1
		dao.UpdateBlogById(id, blog)
	}
	return c.Render(blog, rcnt)
}
