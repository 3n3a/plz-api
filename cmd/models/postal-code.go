package models

type PLZ struct {
	PostalCode int `json:"postal_code" gorm:"primaryKey;column:postal_code"`
	Place string `json:"place" gorm:"column:place"`
}

func (PLZ) TableName() string {
	return "plz"
}