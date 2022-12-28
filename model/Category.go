package model

import (
	"goblog/pkg/e"
	"gorm.io/gorm"
)

// Category 分类
type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	_ = DB.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return e.ErrorCateNameUsed //3001
	}
	return e.SUCCESS
}

// CreateCate 新增分类
func CreateCate(data *Category) int {
	err := DB.Create(&data).Error
	if err != nil {
		return e.ERROR // 500
	}
	return e.SUCCESS
}

// GetCate 查询分类列表
func GetCate(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err = DB.Find(&cate).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	DB.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// EditCate 编辑分类信息
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = DB.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// DeleteCate 删除分类
func DeleteCate(id int) int {
	var cate Category
	err = DB.Where("id = ? ", id).Delete(&cate).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}
