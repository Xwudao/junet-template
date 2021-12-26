package models

//go:generate junet gen db

//gen:qs
type User struct {
	ID       int64  `gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
}
