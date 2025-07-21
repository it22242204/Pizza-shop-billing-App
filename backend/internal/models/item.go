package models

import "gorm.io/gorm"

type Item struct {
    gorm.Model
    Name     string  `gorm:"unique;not null"`
    Category string  `gorm:"not null"`          // pizza | topping | beverage
    Price    int64   `gorm:"not null"`          // cents
    TaxRate  float64 `gorm:"default:0.0"`       // 0.05 = 5 %
}
