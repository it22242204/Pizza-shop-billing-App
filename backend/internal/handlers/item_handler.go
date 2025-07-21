package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/your‑gh‑user/pizza‑backend/internal/services"
    "strconv"
)

type ItemHandler struct {
    svc services.ItemService
}

func NewItemHandler(s services.ItemService) *ItemHandler { return &ItemHandler{s} }

func (h *ItemHandler) Register(r fiber.Router) {
    r.Get("/", h.list)
    r.Post("/", h.create)
    r.Put("/:id", h.update)
    r.Delete("/:id", h.delete)
}

func (h *ItemHandler) list(c *fiber.Ctx) error {
    items, err := h.svc.List()
    if err != nil {
        return fiberErr(c, err)
    }
    return c.JSON(items)
}

func (h *ItemHandler) create(c *fiber.Ctx) error {
    var req services.CreateItemDTO
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    itm, err := h.svc.Create(req)
    if err != nil {
        return fiberErr(c, err)
    }
    return c.Status(201).JSON(itm)
}

func (h *ItemHandler) update(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    var req services.UpdateItemDTO
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    itm, err := h.svc.Update(uint(id), req)
    if err != nil {
        return fiberErr(c, err)
    }
    return c.JSON(itm)
}

func (h *ItemHandler) delete(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    if err := h.svc.Delete(uint(id)); err != nil {
        return fiberErr(c, err)
    }
    return c.SendStatus(204)
}

func fiberErr(c *fiber.Ctx, err error) error {
    return c.Status(500).JSON(fiber.Map{"error": err.Error()})
}
