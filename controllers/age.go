package controllers

import "github.com/gin-gonic/gin"
import . "ExamCode/conf"
import . "ExamCode/models"

// AgeHandler
// @Summary 获取最新age
// @Tags AgeController 年龄相关
// @Accept application/json
// @Produce application/json
// @Router /age [get]
func GetAgeHandler(c *gin.Context) {
	Lock.RLock()
	ret := Age
	defer Lock.RUnlock()
	c.JSON(200, gin.H{
		"age": ret,
	})
}

// AgeHandler
// @Summary 设置最新age
// @Tags AgeController 年龄相关
// @Accept application/json
// @Produce application/json
// @Param object body AgeModel false "请求参数"
// @Router /age [post]
func AddAgeHandler(c *gin.Context) {
	ageModel := AgeModel{}
	err := c.ShouldBind(&ageModel)
	if err != nil || ageModel.Age <= 0 {
		c.JSON(200, gin.H{
			"error": "不合法的参数输入",
		})
		return
	}
	Lock.Lock()
	Age = ageModel.Age
	defer Lock.Unlock()
	c.JSON(200, gin.H{
		"age": ageModel.Age,
	})
}
