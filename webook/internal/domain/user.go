package domain

import "time"

// 领域对象
type User struct{
	Email string
	PassWord string
	Ctime time.Time
}

