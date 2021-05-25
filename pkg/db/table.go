package db

type User struct {
	Id int64 `gorm:"id"`
	Name string `gorm:"name"`
	Sex string `gorm:"sex"`
}

func (User) TableName() string {
	return "user"
}
