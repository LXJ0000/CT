package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() error {
	dsn := "root:root1234@tcp(127.0.0.1:3306)/ct?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn))

	return err
}

//func AddDBData() {
//	db.AutoMigrate(&User{}, &Post{}, &Topic{})
//	users := []User{{Name: "LXJ"}, {Name: "ZN"}, {Name: "HH"}}
//	db.Create(&users)
//	topics := []Topic{
//		{UserId: 1, Title: "first", Content: "The first"},
//		{UserId: 2, Title: "second", Content: "The second"},
//		{UserId: 3, Title: "third", Content: "The third"},
//	}
//	db.Create(&topics)
//	posts := []Post{
//		{UserId: 1, TopicId: 1, Content: "first post"},
//		{UserId: 2, TopicId: 1, Content: "2 post"},
//		{UserId: 2, TopicId: 1, Content: "3 post"},
//		{UserId: 2, TopicId: 2, Content: "4 post"},
//		{UserId: 3, TopicId: 3, Content: "5 post"},
//	}
//	db.Create(&posts)
//}
