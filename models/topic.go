package models

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

type Topic struct {
	Id         int64     `gorm:"column:id"`
	UserId     int64     `gorm:"column:user_id"`
	Title      string    `gorm:"column:title"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
	ViewCount  int       `gorm:"column:count;default:0"`
}

func (Topic) TableName() string {
	return `topic`
}

type TopicDao struct{}

var topicDao *TopicDao
var topicOnce sync.Once

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		},
	)
	return topicDao
}

func (*TopicDao) QueryTopic(id int64) (*Topic, error) {
	var topic Topic
	err := db.Where("id=?", id).Find(&topic).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		//打印日志
		return nil, err
	}
	return &topic, nil
}
