package main

import (
	"fmt"

	"gin_todolist/dao"
	"gin_todolist/routers"
)

func main() {
	// 1.连接 mysql 数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("------connect mysql success!------")
	}
	// 2.设置路由规则
	r := routers.SetupRouter()
	// 3.在指定端口启动
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
} //end main
