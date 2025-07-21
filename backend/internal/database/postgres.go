package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func Connect(dsn string, models ...interface{}) *gorm.DB {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("db connect: %v", err)
    }
    if err := db.AutoMigrate(models...); err != nil {
        log.Fatalf("migrate: %v", err)
    }
    return db
}
