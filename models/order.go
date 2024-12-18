package models

import "time"

type Order struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time
	ProductRef int     `json:"product_id"`
	Product    Product `gorm:"foreignKey:ProductRef"`
	UserRef    int     `json:"user_id"`
	User       User    `gorm:"foreignKey:UserRef"`
}
