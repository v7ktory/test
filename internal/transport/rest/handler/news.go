package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/v7ktory/test/internal/storage/postgres/ref"
)

func (h *Handler) GetList(c *fiber.Ctx) error {
	newsList, err := h.svc.GetList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Success": true,
		"News":    newsList,
	})
}
func (h *Handler) EditPost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("Id"))
	if err != nil || id < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id input")
	}

	news := new(ref.News)
	if err = c.BodyParser(news); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid data input")
	}

	if err = h.svc.Edit(id, *news); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "news wasn't update: "+err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Success": true, "Id": id})
}
