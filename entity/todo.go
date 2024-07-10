package entity

import "gin_todolist/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// TableName : 将 结构体Todo表名设置为t_todolist
func (t Todo) TableName() string {
	return "t_todolist"
}

func CreateTodo(todo *Todo) error {
	return dao.DB.Create(&todo).Error
}

func GetAllTodo() (todoList []Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	return
}

func GetTodoById(id string) (todo *Todo, err error) {
	err = dao.DB.Where("id = ?", id).First(&todo).Error
	return
}

func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteTodoById(id string) (err error) {
	err = dao.DB.Delete(&Todo{}, id).Error
	return
}
