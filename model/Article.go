package model

import (
	"github.com/jinzhu/gorm"
	"goblog/utils/errmsg"
)

type Article struct {
	gorm.Model

	Category Category

	Title   string `gorm:"type:varchar(20);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArticle(data *Article) int {
	// todo 添加用户
	//data.Password = ScryptPw(data.Password)

	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

// todo 查询分类下的所有文章

// todo 查询单个文章

// todo 查询文章列表
func GetArticle(pageSize int, pageNum int) []Article {
	var article []Article
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return article
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
