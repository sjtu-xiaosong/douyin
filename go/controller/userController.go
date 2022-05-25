package controller

import (
	"douyin/go/model"
	"douyin/go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	// 定义一个User变量
	var user model.User
	// 将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&user)
	// 将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err := service.CreateUser(&user)
	// 判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"tip": "测试成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"tip": "测试失败",
		})
	}
}
