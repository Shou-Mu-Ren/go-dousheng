package main

import (
	"douyin/controller"
	"douyin/repository"
	"douyin/util"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := Init(); err != nil {
		os.Exit(-1)
	}

	go controller.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	// if err := r.Run(); err != nil {// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// 	return
	// }

	//改映射可以用localhost
	//window用ip连接监听一般不可以用localhost
	if err := r.Run("192.168.1.103:8080"); err != nil {
		return
	}

}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	if err := util.InitLogger(); err != nil {
		return err
	}
	repository.Pool_Init()
	return nil
}
