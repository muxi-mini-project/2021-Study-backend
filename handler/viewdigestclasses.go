package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func ViewDigestClasses(c *gin.Context) {
	var summaryClasses []model.SummaryClass
	user_id := c.Param("user_id")

	model.DB.Where("user_id = ?", user_id).Find(&summaryClasses)

	c.JSON(200, gin.H{
		"message": "获取成功",
		"data":    summaryClasses,
	})

}