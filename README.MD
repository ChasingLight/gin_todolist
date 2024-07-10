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