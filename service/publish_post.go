package service

import (
	"CT/models"
	"errors"
	"time"
	"unicode/utf8"
)

type PublishPostFlow struct {
	userId  int64
	topicId int64
	content string
	postId  int64
}

func PublishPost(topicId, userId int64, content string) (int64, error) {
	return NewPublishPost(topicId, userId, content).Do()
}

func NewPublishPost(topicId, userId int64, content string) *PublishPostFlow {
	return &PublishPostFlow{
		userId:  userId,
		topicId: topicId,
		content: content,
	}
}

func (f *PublishPostFlow) Do() (int64, error) {

	if err := f.checkParam(); err != nil {
		return 0, err
	}
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.postId, nil
}

func (f *PublishPostFlow) checkParam() error {
	if f.userId <= 0 || f.topicId <= 0 {
		return errors.New("id must be larger than 0")
	}
	if utf8.RuneCountInString(f.content) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

func (f *PublishPostFlow) publish() error {
	post := &models.Post{
		TopicId:    f.topicId,
		UserId:     f.userId,
		Content:    f.content,
		CreateTime: time.Now(),
	}
	if err := models.NewPostDaoInstance().CreatePost(post); err != nil {
		return err
	}
	f.postId = post.Id
	return nil
}
