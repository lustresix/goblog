package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/pkg/e"
	"net/http"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)

	if code == e.SUCCESS {
		code = model.CreateCate(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": e.GetMsg(code),
	})
}

// GetCate 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetCate(pageSize, pageNum)
	code := e.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": e.GetMsg(code),
		},
	)
}

// EditCate 编辑分类名
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == e.SUCCESS {
		model.EditCate(id, &data)
	}
	if code == e.ErrorCateNameUsed {
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": e.GetMsg(code),
		},
	)
}

// DeleteCate 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteCate(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": e.GetMsg(code),
		},
	)
}
