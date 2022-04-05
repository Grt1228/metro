package main

import (
	"fmt"
	"metro/dao"
	"metro/router"
	"metro/settings"
)

func main() {
	//加载配置文件
	fmt.Println("项目启动~~~~~~~~~~~~~~~~~~~")
	fmt.Println("项目加载配置文件~~~~~~~~~~~~~~~~~~~")
	err := settings.Init()
	if err != nil {
		fmt.Printf("项目加在配置文件失败：%s \n ~~~~~~~~~~~~~~~~~~~", err)
	}
	fmt.Println("项目加载配置文件成功~~~~~~~~~~~~~~~~~~~")

	//创建数据库连接
	fmt.Println("加载数据库连接~~~~~~~~~~~~~~~~~~~")
	err = dao.Setup()
	if err != nil {
		fmt.Printf("加载数据库连接失败：%s \n ~~~~~~~~~~~~~~~~~~~", err)
	}
	fmt.Println("加载数据库连接成功~~~~~~~~~~~~~~~~~~~")
	r := router.SetRouters()
	r.Run()
}
