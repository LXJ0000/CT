package router

import (
	"CT/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")

	apiRouter := r.Group("/ct")
	//根据id获取页面信息 /ct/page/1
	//获取页面信息
	apiRouter.GET("/page/:id", func(context *gin.Context) {
		topicId := context.Param("id")
		data := controller.QueryPageInfo(topicId)
		context.JSON(http.StatusOK, data)
	})
	//添加帖子
	apiRouter.POST("post/do", func(context *gin.Context) {
		uid, _ := context.GetPostForm("uid")
		topicId, _ := context.GetPostForm("topic_id")
		content, _ := context.GetPostForm("content")
		data := controller.PublishPost(uid, topicId, content)
		context.JSON(http.StatusOK, data)
	})
}
