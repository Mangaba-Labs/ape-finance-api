package handler

import (
	"strconv"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/services"
	"github.com/gofiber/fiber/v2"
)

// Handler for user service
type Handler struct {
	service services.UserService
}

// CreateUser Handler for POST /user
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var usr = &model.User{}
	if err := c.BodyParser(usr); err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": err})
	}

	newUser, err := h.service.CreateUser(usr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created usr", "data": newUser})
}

//EditUser handler for PUT /user/:id
func (h *Handler) EditUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	var usr model.User
	if err := c.BodyParser(&usr); err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Invalid fields"})
	}

	err = h.service.UpdateUser(&usr, int(id))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "UpdatedUser", "data": usr})
}

// GetUser Handler for GET /user/:id
func (h *Handler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "User ID is invalid"})
	}

	var usr model.User

	usr, err = h.service.GetUserById(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Cannot get user"})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "data": model.Response{
		Email: usr.Email,
		Name:  usr.Name,
	}})
}

// DeleteUser delete user
func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "User ID is invalid"})
	}

	err = h.service.DeleteUser(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Cannot delete user"})
	}

	return c.Status(204).JSON(fiber.Map{"status": "success", "data": nil})
}
