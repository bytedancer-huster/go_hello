package db

type User struct {
	Id int64 `gorm:"id"`
	Name string `gorm:"name"`
	Sex string `gorm:"sex"`
	Password string `gorm:"password"`
}

func (User) TableName() string {
	return "user"
}
