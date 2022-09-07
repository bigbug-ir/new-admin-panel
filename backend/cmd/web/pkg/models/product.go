package models

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	UserRefer   uint    `json:"userRefer"`
	User        User    `gorm:"foreignKey:ID;references:UserRefer"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
