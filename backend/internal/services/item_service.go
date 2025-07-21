package services

import (
    "github.com/your‑gh‑user/pizza‑backend/internal/models"
    "github.com/your‑gh‑user/pizza‑backend/internal/repositories"
)

type ItemService interface {
    List() ([]models.Item, error)
    Create(req CreateItemDTO) (*models.Item, error)
    Update(id uint, req UpdateItemDTO) (*models.Item, error)
    Delete(id uint) error
}

type itemSvc struct{ repo repositories.ItemRepository }

func NewItemService(r repositories.ItemRepository) ItemService {
    return &itemSvc{r}
}

type CreateItemDTO struct {
    Name, Category string
    Price          int64
    TaxRate        float64
}

type UpdateItemDTO = CreateItemDTO

func (s *itemSvc) List() ([]models.Item, error) { return s.repo.All() }

func (s *itemSvc) Create(d CreateItemDTO) (*models.Item, error) {
    itm := &models.Item{
        Name: d.Name, Category: d.Category,
        Price: d.Price, TaxRate: d.TaxRate,
    }
    return itm, s.repo.Create(itm)
}

func (s *itemSvc) Update(id uint, d UpdateItemDTO) (*models.Item, error) {
    itm, err := s.repo.GetByID(id)
    if err != nil {
        return nil, err
    }
    itm.Name, itm.Category, itm.Price, itm.TaxRate =
        d.Name, d.Category, d.Price, d.TaxRate
    return itm, s.repo.Update(itm)
}

func (s *itemSvc) Delete(id uint) error { return s.repo.Delete(id) }
