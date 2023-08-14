package models

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

type Post struct {
	Id         int64     `gorm:"column:id"`
	UserId     int64     `gorm:"column:user_id"`
	TopicId    int64     `gorm:"column:topic_id"`
	Content    string    `gorm:"column:content"`
	DiggCount  int64     `gorm:"column:digg_count,default:0"` // 点赞数
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Post) TableName() string {
	return `post`
}

type PostDao struct{}

var postDao *PostDao
var postOnce sync.Once

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) CreatePost(post *Post) error {
	if err := db.Create(post).Error; err != nil {
		//	对应日志
		return err
	}
	return nil
}

func (*PostDao) QueryPost(id int64) (*Post, error) {
	var post Post
	err := db.Where("id=?", id).Find(&post).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		//
		return nil, err
	}
	return &post, nil
}
func (*PostDao) QueryPostList(topicId int64) ([]*Post, error) {
	var posts []*Post
	err := db.Where("topic_id=?", topicId).Find(&posts).Error
	if err != nil {
		//
		return nil, err
	}
	return posts, nil
}
