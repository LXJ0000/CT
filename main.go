package main

import (
	"CT/models"
	"CT/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := models.InitDatabase(); err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	r := gin.Default()

	router.InitRouter(r)

	if err := r.Run(); err != nil {
		return
	}
}
