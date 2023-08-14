package models

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

type User struct {
	Id         int64     `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	CreateTime time.Time `gorm:"column:create_time"`
	ModifyTime time.Time `gorm:"column:modify_time"`
}

// TableName 定义数据库表明 默认为users
func (User) TableName() string {
	return `user`
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once //用来确保`NewUserDaoInstance()`函数只被执行一次
//使用`sync.Once`可以避免重复创建和管理多个实例的问题，确保只有一个实例被创建并被多个goroutine共享。这样可以提高代码的效率和性能。

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryUser(id int64) (*User, error) {
	var user User
	err := db.Where("id=?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		//打印日志
		return nil, err
	}
	return &user, nil
}

func (*UserDao) QueryUserMap(ids []int64) (map[int64]*User, error) {
	var users []*User
	err := db.Where("id in (?)", ids).Find(&users).Error
	if err != nil {
		//打印日志
		return nil, err
	}
	userMap := make(map[int64]*User)
	for _, user := range users {
		userMap[user.Id] = user
	}

	return userMap, nil
}
