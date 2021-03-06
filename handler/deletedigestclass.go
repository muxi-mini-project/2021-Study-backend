package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

// @Summary 删除书摘的分类
// @Description "删除我的书摘分类里的类别"
// @Tags digest
// @Accept json
// @Produce json
// @Success 200 "成功删除"
// @Router /digest/mysummary/:user_id/clasees_edit [delete]
func DeleteDigestClass(c *gin.Context) {
	var summaryClass model.SummaryClass
	var summaries []model.Summary
	class_id := c.Query("class_id")

	model.DB.Where("class_id = ?", class_id).Delete(&summaryClass)
	model.DB.Where("class_id = ?", class_id).Delete(&summaries)

	c.JSON(200, gin.H{
		"message": "成功删除",
	})

}