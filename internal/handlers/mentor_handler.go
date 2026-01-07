package handlers

import (
	"project-kelas-santai/internal/models"
	"project-kelas-santai/internal/services"
	"project-kelas-santai/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MentorHandler struct {
	service services.MentorService
}

func NewMentorHandler(service services.MentorService) *MentorHandler {
	return &MentorHandler{
		service: service,
	}
}

func (h *MentorHandler) CreateMentor(c *fiber.Ctx) error {
	var mentor models.Mentor
	if err := c.BodyParser(&mentor); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := h.service.CreateMentor(&mentor); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to create mentor", err.Error())
	}

	return utils.SendSuccess(c, "Mentor created successfully", mentor)
}

func (h *MentorHandler) GetAllMentors(c *fiber.Ctx) error {
	mentors, err := h.service.GetAllMentors()
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch mentors", err.Error())
	}
	return utils.SendSuccess(c, "Mentors fetched successfully", mentors)
}

func (h *MentorHandler) GetMentorByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid mentor ID", nil)
	}

	mentor, err := h.service.GetMentorByID(id)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, "Mentor not found", err.Error())
	}

	return utils.SendSuccess(c, "Mentor fetched successfully", mentor)
}

func (h *MentorHandler) UpdateMentor(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid mentor ID", nil)
	}

	var mentor models.Mentor
	if err := c.BodyParser(&mentor); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	mentor.ID = id
	if err := h.service.UpdateMentor(&mentor); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to update mentor", err.Error())
	}

	return utils.SendSuccess(c, "Mentor updated successfully", nil)
}

func (h *MentorHandler) DeleteMentor(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid mentor ID", nil)
	}

	if err := h.service.DeleteMentor(id); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to delete mentor", err.Error())
	}

	return utils.SendSuccess(c, "Mentor deleted successfully", nil)
}
