package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/pkg/e"
	"goblog/utils"
	"net/http"
	"time"
)

// UserLogin 后台登陆
func UserLogin(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int

	code = model.CheckLogin(formData.Username, formData.Password)

	if code == e.SUCCESS {
		setToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": e.GetMsg(code),
			"token":   token,
		})
	}

}

// token生成函数
func setToken(c *gin.Context, user model.User) {
	j := utils.NewJWT()
	claims := utils.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "GoBlog",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  e.ERROR,
			"message": e.GetMsg(e.ERROR),
			"token":   token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    user.Username,
		"id":      user.ID,
		"message": e.GetMsg(200),
		"token":   token,
	})
	return
}
