package controller

// import规范：标准库包、程序内部包、第三方包
import (
	"fmt"
	"net/http"

	"gin_todolist/entity"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	var todo entity.Todo
	c.BindJSON(&todo)
	// 调用实体层保存
	if err := entity.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetAllTodo(c *gin.Context) {
	todoList, err := entity.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	// 1.获取请求：路径参数-id、Json参数status
	id := c.Param("id")
	var req entity.Todo
	c.BindJSON(&req)
	newStatus := req.Status
	// 2.根据id从数据库中查询记录
	dbTodo, err := entity.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "未找到指定id的记录" + err.Error(),
		})
		return
	}
	// 3.更新status
	dbTodo.Status = newStatus
	if err := entity.UpdateTodo(dbTodo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, dbTodo)
}

func DeleteTodo(c *gin.Context) {
	// 1.获取请求路径参数id
	id := c.Param("id")
	// 2.判断记录是否存在
	_, err := entity.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": fmt.Sprintf("未找到id:%s 的记录", id),
		})
		return
	}
	// 3.执行删除
	if err := entity.DeleteTodoById(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": id + " deleted",
	})
}
