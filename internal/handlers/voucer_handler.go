package handlers

import (
	"project-kelas-santai/internal/models"
	"project-kelas-santai/internal/services"
	"project-kelas-santai/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type VoucerHandler struct {
	service services.VoucerService
}

func NewVoucerHandler(service services.VoucerService) *VoucerHandler {
	return &VoucerHandler{
		service: service,
	}
}

// func (h *VoucerHandler) UseVoucer(c *fiber.Ctx) error {
// 	type Request struct {
// 		VoucerId string `json:"voucer_id"`
// 		Id
// 	}
// 	var req Request
// 	if err := c.BodyParser(&req); err != nil {
// 		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
// 	}

// 	if err := h.service.UseVoucer(req.VoucerId); err != nil {
// 		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to use voucer", err.Error())
// 	}

// 	return utils.SendSuccess(c, "Voucer used successfully", nil)
// }

func (h *VoucerHandler) CreateVoucer(c *fiber.Ctx) error {
	var voucer models.Voucer
	if err := c.BodyParser(&voucer); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	voucer.ID = uuid.New()

	if err := h.service.CreateVoucer(&voucer); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to create voucer", err.Error())
	}

	return utils.SendSuccess(c, "Voucer created successfully", voucer)
}

func (h *VoucerHandler) GetAllVoucer(c *fiber.Ctx) error {
	voucers, err := h.service.GetVoucerAll()
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch voucers", err.Error())
	}

	return utils.SendSuccess(c, "Voucers fetched successfully", voucers)
}

func (h *VoucerHandler) GetVoucerById(c *fiber.Ctx) error {
	id := c.Params("id")
	idVoucer, _ := uuid.Parse(id)
	voucer, err := h.service.GetVoucerById(idVoucer)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, "Voucer not found", err.Error())
	}

	return utils.SendSuccess(c, "Voucer fetched successfully", voucer)
}

func (h *VoucerHandler) UpdateVoucer(c *fiber.Ctx) error {
	id := c.Params("id")
	idVoucer, _ := uuid.Parse(id)
	var voucer models.Voucer
	if err := c.BodyParser(&voucer); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	// Fetch existing voucer to ensure it exists
	existingVoucer, err := h.service.GetVoucerById(idVoucer)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, "Voucer not found", err.Error())
	}

	// Update fields
	existingVoucer.Name = voucer.Name
	existingVoucer.Discount = voucer.Discount
	// Add other fields updates as necessary

	if err := h.service.UpdateVoucer(existingVoucer); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to update voucer", err.Error())
	}

	return utils.SendSuccess(c, "Voucer updated successfully", existingVoucer)
}

func (h *VoucerHandler) DeleteVoucer(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeleteVoucer(id); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to delete voucer", err.Error())
	}

	return utils.SendSuccess(c, "Voucer deleted successfully", nil)
}
