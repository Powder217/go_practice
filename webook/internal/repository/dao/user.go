package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)
var(
	ErrUserDuplicateEmail=errors.New("邮箱冲突")
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
	err:= dao.db.WithContext(ctx).Create(&u).Error
	
	if MysqlErr,ok:=err.(*mysql.MySQLError);ok {
		const uniqueConflictsErrNo uint16=1062
		if MysqlErr.Number==uniqueConflictsErrNo{
			return ErrUserDuplicateEmail
		}
	}
	return err
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