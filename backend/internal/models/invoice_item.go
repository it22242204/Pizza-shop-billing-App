package models

import "gorm.io/gorm"

type InvoiceItem struct {
    gorm.Model
    InvoiceID uint
    ItemID    uint
    Qty       int64
    UnitPrice int64
    LineTotal int64
    Item      Item `gorm:"foreignKey:ItemID"`
}
