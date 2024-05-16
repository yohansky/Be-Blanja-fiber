package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       uint `json:"id"`
	Name     string
	Email    string
	Phone    string
	Store    string
	Password []byte `json:"-"`
	RoleId   uint
	Role     Role `gorm:"foreignKey:RoleId"`
	Photo    string
	// StoreDesc string
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

func (user *User) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&User{}).Count(&total)

	return total
}

func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User

	db.Preload("Role").Offset(offset).Limit(limit).Find(&users)

	return users
}
