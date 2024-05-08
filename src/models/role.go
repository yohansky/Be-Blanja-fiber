package models

type Role struct {
	Id          uint `json:"id"`
	Name        string
	Permissions []Permission `gorm:"many2many:role_permissions"`
}
