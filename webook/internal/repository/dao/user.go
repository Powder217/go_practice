package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserDAO struct{
	db *gorm.DB
}


func NewUserDAO(db *gorm.DB)*UserDAO{
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context,u User)error{
	now:=time.Now().UnixMilli()
	u.Uptime=now
	u.Ctime=now
	return dao.db.WithContext(ctx).Create(&u).Error
}

// 直接对应数据库表结构
type User struct{
	Id int64 `gorm:"primaryKey,autoIncrement"`
	Email string `gorm:"unique"`
	PassWord string
	// 创建时间 毫秒数
	Ctime int64 
	// 更新时间 毫秒数
	Uptime int64
}