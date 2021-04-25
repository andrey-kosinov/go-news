package database

type User struct {
	ID       uint     `gorm:"primaryKey"`
  	Login    string
  	Password string
}

type News struct {
	ID    uint     `gorm:"primaryKey"`
	Title string
	Body  string
}

type Session struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string
	UserId     uint
	TimeToKill int64
	User User
}