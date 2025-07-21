package models

import "gorm.io/gorm"

type Invoice struct {
    gorm.Model
    CustomerName  string         `gorm:"not null"`
    Subtotal      int64          // cents
    Tax           int64
    Total         int64
    Items         []InvoiceItem
}
