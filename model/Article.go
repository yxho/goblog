package model

import (
	"github.com/jinzhu/gorm"
	"goblog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(20);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

// 查询分类下的所有文章
func GetCategoryArticle(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArtList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST_USED
	}
	return cateArtList,errmsg.SUCCES
}

// 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ART_NOT_EXIST
	}
	return article, errmsg.SUCCES

}

// 查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCES
}

// 编辑用户信息
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	db.Model(&article).Where("id = ?", id).Updates(maps)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES

}

// 删除文章
func DeleteArticle(id int) int {
	var article Article

	err = db.Where("id=?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}
