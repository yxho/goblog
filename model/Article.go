package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model

	Category Category

	Title   string `gorm:"type:varchar(20);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}
