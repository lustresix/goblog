package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/pkg/e"
	"goblog/utils/validator"
	"net/http"
	"strconv"
)

var code int

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	_ = c.ShouldBindJSON(&data)
	// 验证规范
	msg, code = validator.Validate(&data)
	if code != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}

	if code = model.CheckUser(data.Username); code == e.SUCCESS {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": e.GetMsg(code),
	})

}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNumber, _ := strconv.Atoi(c.Query("pagenum"))
	//username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNumber == 0 {
		pageNumber = 1
	}

	data := model.GetUsers(pageSize, pageNumber)
	code := e.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": e.GetMsg(code),
		},
	)

}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"total":   1,
			"message": e.GetMsg(code),
		},
	)
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.CheckUpUser(id, data.Username)
	if code == e.SUCCESS {
		model.EditUser(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": e.GetMsg(code),
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": e.GetMsg(code),
	})
}
