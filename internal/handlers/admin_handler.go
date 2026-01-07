package handlers

import (
	"project-kelas-santai/internal/models"
	"project-kelas-santai/internal/services"
	"project-kelas-santai/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AdminHandler struct {
	service services.AdminService
}

func NewAdminHandler(service services.AdminService) *AdminHandler {
	return &AdminHandler{
		service: service,
	}
}

func (h *AdminHandler) CreateAdmin(c *fiber.Ctx) error {
	var admin models.Admin
	if err := c.BodyParser(&admin); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := h.service.CreateAdmin(&admin); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to create admin", err.Error())
	}

	return utils.SendSuccess(c, "Admin created successfully", admin)
}

func (h *AdminHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var request LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	token, err := h.service.Login(request.Email, request.Password)
	if err != nil {
		return utils.SendError(c, fiber.StatusUnauthorized, "Login failed", err.Error())
	}

	return utils.SendSuccess(c, "Login successful", fiber.Map{
		"token": token,
	})
}

func (h *AdminHandler) GetAdminByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid admin ID", nil)
	}

	admin, err := h.service.GetAdminByID(id)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, "Admin not found", err.Error())
	}

	return utils.SendSuccess(c, "Admin fetched successfully", admin)
}

func (h *AdminHandler) UpdateAdmin(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid admin ID", nil)
	}

	var admin models.Admin
	if err := c.BodyParser(&admin); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	admin.ID = id
	if err := h.service.UpdateAdmin(&admin); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to update admin", err.Error())
	}

	return utils.SendSuccess(c, "Admin updated successfully", nil)
}

func (h *AdminHandler) DeleteAdmin(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid admin ID", nil)
	}

	if err := h.service.DeleteAdmin(id); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to delete admin", err.Error())
	}

	return utils.SendSuccess(c, "Admin deleted successfully", nil)
}
