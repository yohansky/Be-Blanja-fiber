package models

type Payment struct {
	Id           uint
	UserId       uint
	User         User `gorm:"foreignKey:UserId"`
	BankId       uint
	Bank         Bank `gorm:"foreignKey:BankId"`
	TotalPayment string
	Status       string
}
