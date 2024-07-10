package routers

import (
	"gin_todolist/controller"
	
	"github.com/gin-gonic/gin"
)

func SetupRouter() (r *gin.Engine) {
	r = gin.Default()

	// 加载静态资源（css、js）
	r.Static("/static", "./static")
	// 加载模版文件
	r.LoadHTMLGlob("templates/*")

	// 单个路由：主页
	r.GET("/", controller.IndexHandler)
	// 路由组：待办事项增删改查
	v1Group := r.Group("v1")
	{
		// 新增
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有
		v1Group.GET("/todo", controller.GetAllTodo)
		// 更新
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return
}
