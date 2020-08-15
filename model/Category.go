package model

import (
	"github.com/jinzhu/gorm"
	"goblog/utils/errmsg"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED // 2001
	}
	return errmsg.SUCCES
}

// 新增分类
func CreateCategory(data *Category) int {
	//data.Password = ScryptPw(data.Password)

	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

// todo 查询分类下的所有文章

// 查询分类列表
func GetCategory(pageSize int, pageNum int) ([]Category, int) {
	var cate []Category
	var total int
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// 编辑用户信息
func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	db.Model(&cate).Where("id = ?", id).Updates(maps)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES

}

// 删除用户
func DeleteCategory(id int) int {
	var cate Category

	err = db.Where("id=?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}
