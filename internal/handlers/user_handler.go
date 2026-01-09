package handlers

import (
	//"project-kelas-santai/internal/database"

	"project-kelas-santai/internal/models"
	"project-kelas-santai/internal/services"
	"project-kelas-santai/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	type UserDto struct {
		ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
		Name      string    `json:"name"`
		Email     string    `gorm:"uniqueIndex" json:"email"`
		NoTelp    string    `json:"no_telp"`
		Address   string    `json:"address"`
		City      string    `json:"city"`
		Province  string    `json:"province"`
		Gender    string    `json:"gender"`
		Password  string    `json:"password"` // Don't return password in JSON
		CreatedAt string    `json:"created_at"`
		UpdatedAt string    `json:"updated_at"`
	}

	var user UserDto
	if err := c.BodyParser(&user); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		NoTelp:   user.NoTelp,
		Address:  user.Address,
		City:     user.City,
		Province: user.Province,
		Gender:   user.Gender,
		Password: user.Password,
	}

	if err := h.service.CreateUser(&newUser); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to create user", err.Error())
	}

	//database.DB.Create(&user)
	go func() {
		utils.SendOrderSuccessEmail(user.Email, user.Name, "Selamat datang di Kelas Santai")
	}()

	return utils.SendSuccess(c, "User created successfully", nil)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
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

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch users", err.Error())
	}

	return utils.SendSuccess(c, "Users fetched successfully", users)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	userIDStr := c.Locals("user_id").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, "User not found", err.Error())
	}

	return utils.SendSuccess(c, "User fetched successfully", user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userIDStr := c.Locals("user_id").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := h.service.UpdateUser(userID, &user); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to update user", err.Error())
	}

	return utils.SendSuccess(c, "User updated successfully", nil)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}

	if err := h.service.DeleteUser(id); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to delete user", err.Error())
	}

	return utils.SendSuccess(c, "User deleted successfully", nil)
}
