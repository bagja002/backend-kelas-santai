package handlers

import (
	"project-kelas-santai/internal/services"
	"project-kelas-santai/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserCourseHandler struct {
	service services.UserCourseService
}

type TransactionHandler struct {
	service services.TransactionService
}

func NewUserCourseHandler(service services.UserCourseService) *UserCourseHandler {
	return &UserCourseHandler{
		service: service,
	}
}

func NewTransactionHandler(service services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

func (h *UserCourseHandler) EnrollCourse(c *fiber.Ctx) error {
	type EnrollRequest struct {
		CourseID string `json:"course_id"`
	}

	var request EnrollRequest
	if err := c.BodyParser(&request); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	userIDStr := c.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)
	courseID, err := uuid.Parse(request.CourseID)

	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid UUIDs", err.Error())
	}

	if err := h.service.EnrollCourse(userID, courseID); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to enroll in course", err.Error())
	}

	return utils.SendSuccess(c, "Enrolled successfully", nil)
}

func (h *TransactionHandler) PaymentCourse(c *fiber.Ctx) error {

	type PaymentRequest struct {
		CourseID string `json:"course_id"`
	}

	var request []PaymentRequest
	if err := c.BodyParser(&request); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	userIDStr := c.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)
	ListCourseId := []uuid.UUID{}
	for _, course := range request {
		courseID, err := uuid.Parse(course.CourseID)
		if err != nil {
			return utils.SendError(c, fiber.StatusBadRequest, "Invalid UUIDs", err.Error())
		}
		ListCourseId = append(ListCourseId, courseID)
	}

	err, directURL := h.service.PaymentCourse(userID, ListCourseId)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to enroll in course", err.Error())
	}

	var Link struct {
		DirectURL string `json:"direct_url"`
	}

	Link.DirectURL = directURL

	return utils.SendSuccess(c, "Payment successful", Link)
}

func (h *UserCourseHandler) GetMyCourses(c *fiber.Ctx) error {
	userIDStr := c.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	userCourses, err := h.service.GetUserCourses(userID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch user courses", err.Error())
	}

	return utils.SendSuccess(c, "User courses fetched successfully", userCourses)
}

func (h *UserCourseHandler) GetCoursePending(c *fiber.Ctx) error {
	courseIDStr := c.Locals("user_id").(string)
	courseID, _ := uuid.Parse(courseIDStr)

	courseDetail, err := h.service.GetCoursePending(courseID, "pending")
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch course detail", err.Error())
	}

	return utils.SendSuccess(c, "Course detail fetched successfully", courseDetail)
}

func (h *UserCourseHandler) DeleteCourse(c *fiber.Ctx) error {
	userIDStr := c.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	courseIDStr := c.Query("course_id")
	courseID, _ := uuid.Parse(courseIDStr)

	courseDetail := h.service.Delete(userID, courseID)
	if courseDetail != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to delete course", courseDetail.Error())
	}

	return utils.SendSuccess(c, "Course deleted successfully", nil)
}
