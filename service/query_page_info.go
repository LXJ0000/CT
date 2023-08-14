package service

import (
	"CT/models"
	"errors"
	"fmt"
	"sync"
)

type TopicInfo struct {
	Topic *models.Topic
	User  *models.User
}

type PostInfo struct {
	Post *models.Post
	User *models.User
}

type PageInfo struct {
	TopicInfo *TopicInfo
	PostList  []*PostInfo
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo
	topic    *models.Topic
	posts    []*models.Post
	userMap  map[int64]*models.User
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfo(topicId).Do()
}

func NewQueryPageInfo(topicId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topicId,
	}
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New(`topic id must be larger than 0`)
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)
	var topicErr, postErr error
	go func() {
		defer wg.Done()
		topic, err := models.NewTopicDaoInstance().QueryTopic(f.topicId)
		if err != nil {
			topicErr = err
			return
		}
		f.topic = topic
	}()
	go func() {
		defer wg.Done() // 保证最后一定执行

		posts, err := models.NewPostDaoInstance().QueryPostList(f.topicId)
		if err != nil {
			postErr = err
			return
		}
		f.posts = posts
	}()
	wg.Wait()
	if topicErr != nil {
		return topicErr
	}
	if postErr != nil {
		return postErr
	}

	//获取用户信息
	uids := []int64{f.topic.Id}
	for _, post := range f.posts {
		uids = append(uids, post.Id)
	}
	userMap, err := models.NewUserDaoInstance().QueryUserMap(uids)
	if err != nil {
		return err
	}
	f.userMap = userMap
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	userMap := f.userMap
	topicUser, ok := userMap[f.topic.UserId]
	if !ok {
		return errors.New(`has no topic user info`)
	}
	postList := make([]*PostInfo, 0)
	for _, post := range f.posts {
		PostUser, ok := userMap[post.UserId]
		if !ok {
			return errors.New("has no post user info for " + fmt.Sprint(post.UserId))
		}
		postList = append(postList, &PostInfo{
			Post: post,
			User: PostUser,
		})
	}
	f.pageInfo = &PageInfo{
		TopicInfo: &TopicInfo{
			Topic: f.topic,
			User:  topicUser,
		},
		PostList: postList,
	}
	return nil
}
