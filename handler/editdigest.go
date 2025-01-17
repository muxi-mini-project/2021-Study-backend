﻿package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

// @Summary "编辑书摘"
// @Description "通过编辑修改我发布的书摘"
// @Tags digest
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param summary_id path string true "summary_id"
// @Param summaryInfo body model.SummaryInfo true "summaryInfo"
// @Success 200 {object} []model.Summary "编辑成功"
// @Failure 404 "编辑失败，数据为空"
// @Failure 400 "编辑错误"
// @Router /digest/person/{summary_id} [put]
func EditDigest(c *gin.Context) {

	token := c.Request.Header.Get("token")
	_, err0 := model.VerifyToken(token)
	if err0 != nil {
		c.JSON(404, gin.H{"message": "认证失败"})
		return
	}

	var summary model.Summary
	var summaryInfo model.SummaryInfo
	var book model.Book
	summary_id := c.Param("summary_id")

	model.DB.First(&summary, summary_id)

	if summary.Id == 0 {
		c.JSON(404, gin.H{
			"message": "编辑失败，无数据",
		})
		return
	}

	err := c.BindJSON(&summaryInfo)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "编辑错误",
		})
		return
	}

	model.DB.Model(&summary).Update(map[string]interface{}{
		"title":               summaryInfo.Title,
		"chapter":             summaryInfo.Chapter,
		"summary_information": summaryInfo.Summary_information,
		"thought":             summaryInfo.Thought,
		"public":              summaryInfo.Public,
	})

	model.DB.Where("book_name = ?", summary.Title).First(&book)
	/*
		if book.Book_id == 0 {
			book.Book_id = 1
		}
	*/

	model.DB.Model(&summary).Update("book_id", book.Book_id)

	c.JSON(200, gin.H{
		"message": "编辑成功",
		"data":    summary,
	})

}
