package services_test

import (
    "testing"
    "github.com/your‑gh‑user/pizza‑backend/internal/services"
    "github.com/stretchr/testify/require"
)

func TestPriceValidation(t *testing.T) {
    // example: enforce price > 0
    svc := &services.ItemServiceMock{} // create a tiny mock
    _, err := svc.Create(services.CreateItemDTO{Price: -10})
    require.Error(t, err)
}
