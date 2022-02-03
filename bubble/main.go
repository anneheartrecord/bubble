package main

import (
	_ "github.com/go-sql-driver/mysql"
	"golangstudy/gowebjinjie/bubble/dao"
	"golangstudy/gowebjinjie/bubble/routers"
)

//初始化数据库




func main()  {
	err:=dao.InitDB()
	if err!=nil {
		panic(err.Error())
	}
	//加载静态文件和html文件
	r:=routers.SetupRouters()
	r.Run(":8080")
}