package model

import (
	"goblog/pkg/e"
	"gorm.io/gorm"
)

// Article 文章
type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

// CreateArt 添加文章
func CreateArt(data *Article) int {
	err := DB.Create(&data).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// GetCateArt 查询分类下的文章
func GetCateArt(id, pageSize, pageNum int) ([]Article, int, int64) {
	var ArtList []Article
	var total int64
	err = DB.Preload("Category").Limit(pageSize).
		Offset((pageNum-1)*pageSize).Where("cid =?", id).Find(&ArtList).Error
	DB.Model(&ArtList).Where("cid =?", id).Count(&total)
	if err != nil {
		return ArtList, e.ErrorArtNotExist, 0
	}
	return ArtList, e.SUCCESS, total
}

// GetArt 查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64

	err = DB.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("Category").Find(&articleList).Error
	// 单独计数
	DB.Model(&articleList).Count(&total)
	if err != nil {
		return nil, e.ERROR, 0
	}
	return articleList, e.SUCCESS, total

}

// EditArt 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = DB.Model(&art).Where("id = ? ", id).Updates(&maps).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// DeleteArt 删除文章
func DeleteArt(id int) int {
	var art Article
	err := DB.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// GetArtInfo 查询单个文章信息
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := DB.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, e.ErrorArtNotExist // 2001
	}
	return art, e.SUCCESS
}
