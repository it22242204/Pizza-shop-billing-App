package repositories

import (
    "github.com/your‑gh‑user/pizza‑backend/internal/models"
    "gorm.io/gorm"
)

type ItemRepository interface {
    All() ([]models.Item, error)
    Create(item *models.Item) error
    Update(item *models.Item) error
    Delete(id uint) error
    GetByID(id uint) (*models.Item, error)
}

type itemRepo struct{ db *gorm.DB }

func NewItemRepository(db *gorm.DB) ItemRepository {
    return &itemRepo{db}
}

func (r *itemRepo) All() ([]models.Item, error) {
    var items []models.Item
    return items, r.db.Find(&items).Error
}

func (r *itemRepo) Create(i *models.Item) error  { return r.db.Create(i).Error }
func (r *itemRepo) Update(i *models.Item) error  { return r.db.Save(i).Error }
func (r *itemRepo) Delete(id uint) error         { return r.db.Delete(&models.Item{}, id).Error }
func (r *itemRepo) GetByID(id uint) (*models.Item, error) {
    var i models.Item
    err := r.db.First(&i, id).Error
    return &i, err
}
