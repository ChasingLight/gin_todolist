1. 项目初始化，执行命令
```
# 使用 go mod 进行项目包管理
go mod init gin_todolist
# 下载并引入对应依赖（gin框架、gorm框架、mysql驱动）
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

2. 待办事项的增删改查，gin web demo。

3. 资料参考
视频地址：https://www.bilibili.com/video/BV1gJ411p7xC
对应笔记：https://gitee.com/moxi159753/LearningNotes/tree/master/Golang/Gin%E6%A1%86%E6%9E%B6/1_Gin%E5%86%85%E5%AE%B9%E4%BB%8B%E7%BB%8D