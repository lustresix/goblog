package model

import (
	"encoding/base64"
	"goblog/pkg/e"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

// User 用户表
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// CheckUser 查询用户是否存在
func CheckUser(username string) int {
	var user User
	_ = DB.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return e.ErrorUsernameUsed //1001
	}
	return e.SUCCESS // 200
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPwd(data.Password)
	err := DB.Create(&data).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	// 分页
	err := DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// EditUser 编辑用户--密码以外的
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = DB.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = DB.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// ScryptPwd 加密
func ScryptPwd(password string) string {
	const PwdHashByte = 10
	salt := make([]byte, 8)
	salt = []byte{200, 20, 9, 29, 15, 50, 80, 7}

	key, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, PwdHashByte)
	if err != nil {
		log.Fatal(err)
	}
	FinPwd := base64.StdEncoding.EncodeToString(key)
	return FinPwd
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	DB.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return e.SUCCESS
	}
	if user.ID > 0 {
		return e.ErrorUsernameUsed //1001
	}
	return e.SUCCESS
}

// GetUser 获取单个用户
func GetUser(id int) (User, int) {
	var user User
	err := DB.Limit(1).Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, e.SUCCESS
	}
	return user, e.ERROR
}

// CheckLogin 登录认证
func CheckLogin(username, password string) int {
	var user User
	DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return e.ErrorUserNotExist // 1003
	}
	if user.Role != 1 {
		return e.ErrorUserNoRight // 1008
	}
	if user.Password != ScryptPwd(password) {
		return e.ErrorPasswordWrong // 1002
	}
	return e.SUCCESS
}
