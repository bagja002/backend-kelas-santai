package handlers

import (
	"project-kelas-santai/internal/models"
	"project-kelas-santai/internal/services"
	"project-kelas-santai/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CourseHandler struct {
	services services.CourseService
}

func NewCourseHandler(services services.CourseService) CourseHandler {
	return CourseHandler{
		services: services,
	}
}

func (h *CourseHandler) UploadFile(c *fiber.Ctx) error {
	file, _ := c.FormFile("file")
	if file != nil {
		filePath, err := utils.HandleSingleFileUpload(c, "file", "public/uploads/courses")
		if err != nil {
			return utils.SendError(c, fiber.StatusInternalServerError, "Failed to upload file", err.Error())
		}
		return utils.SendSuccess(c, "File uploaded successfully", filePath)
	}
	return utils.SendError(c, fiber.StatusBadRequest, "No file uploaded", nil)
}

func (h *CourseHandler) CreateCourse(c *fiber.Ctx) error {
	var course models.Course
	if err := c.BodyParser(&course); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	// Handle Image Upload
	file, _ := c.FormFile("picture")
	if file != nil {
		picturePath, err := utils.HandleSingleFileUpload(c, "picture", "public/uploads/courses")
		if err != nil {
			return utils.SendError(c, fiber.StatusInternalServerError, "Failed to upload image", err.Error())
		}
		course.Picture = picturePath
	}

	if err := h.services.CreateCourse(&course); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to create course", err.Error())
	}

	return utils.SendSuccess(c, "Course created successfully", course)

}

func (h *CourseHandler) GetAllCourse(c *fiber.Ctx) error {
	category := c.Query("category")
	status := c.Query("status")
	courses, err := h.services.GetAllCourse(category, status)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch courses", err.Error())
	}
	return utils.SendSuccess(c, "Courses fetched successfully", courses)
}

func (h *CourseHandler) GetCourseByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid course ID", nil)
	}
	course, err := h.services.GetCourseByID(id)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch course", err.Error())
	}
	return utils.SendSuccess(c, "Course fetched successfully", course)
}

func (h *CourseHandler) UpdateCourse(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid course ID", nil)
	}
	var course models.Course
	if err := c.BodyParser(&course); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	course.ID = id
	file, _ := c.FormFile("picture")
	if file != nil {
		picturePath, err := utils.HandleSingleFileUpload(c, "picture", "public/uploads/courses")
		if err != nil {
			return utils.SendError(c, fiber.StatusInternalServerError, "Failed to upload image", err.Error())
		}
		course.Picture = picturePath
	}

	if err := h.services.UpdateCourse(&course); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to update course", err.Error())
	}
	return utils.SendSuccess(c, "Course updated successfully", course)
}

func (h *CourseHandler) DeleteCourse(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid course ID", nil)
	}
	if err := h.services.DeleteCourse(id); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to delete course", err.Error())
	}
	return utils.SendSuccess(c, "Course deleted successfully", nil)
}
