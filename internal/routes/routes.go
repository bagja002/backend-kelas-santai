package routes

import (
	"project-kelas-santai/internal/config"
	"project-kelas-santai/internal/database"
	"project-kelas-santai/internal/handlers"
	"project-kelas-santai/internal/middleware"
	"project-kelas-santai/internal/models"
	"project-kelas-santai/internal/repository"
	"project-kelas-santai/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//ok

func SetupRoutes(app *fiber.App, cfg *config.Config) {
	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Health Check
	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// Dependency Injection
	// Dependency Injection & Migration
	database.DB.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Mentor{},
		&models.Course{},
		&models.Curiculum{},
		&models.UserCourse{},
		&models.Transaction{},
		&models.DetailTransaction{},
	)

	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	courseRepo := repository.NewCourseRepository(cfg)
	courseService := services.NewCourseService(courseRepo)
	courseHandler := handlers.NewCourseHandler(courseService)

	adminRepo := repository.NewAdminRepository()
	adminService := services.NewAdminService(adminRepo)
	adminHandler := handlers.NewAdminHandler(adminService)

	mentorRepo := repository.NewMentorRepository()
	mentorService := services.NewMentorService(mentorRepo)
	mentorHandler := handlers.NewMentorHandler(mentorService)

	userCourseRepo := repository.NewUserCourseRepository()
	userCourseService := services.NewUserCourseService(userCourseRepo, cfg)
	userCourseHandler := handlers.NewUserCourseHandler(userCourseService)

	transactionRepo := repository.NewTransactionRepository()
	transactionService := services.NewTransactionService(transactionRepo, cfg)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// User Routes
	users := v1.Group("/users")
	users.Post("/register", userHandler.CreateUser) // Public: Register
	users.Post("/login", userHandler.Login)         // Public: Login

	usersProtected := users.Group("/")
	usersProtected.Use(middleware.Protected())
	usersProtected.Get("/", userHandler.GetAllUsers)
	usersProtected.Get("/getById", userHandler.GetUserByID)
	usersProtected.Put("/update", userHandler.UpdateUser)
	usersProtected.Delete("/:id", userHandler.DeleteUser)

	// Admin Routes
	admins := v1.Group("/admins")
	admins.Post("/register", adminHandler.CreateAdmin) // Setup only
	admins.Post("/login", adminHandler.Login)

	adminsProtected := admins.Group("/")
	adminsProtected.Use(middleware.Protected(), middleware.AdminProtected())
	adminsProtected.Get("/:id", adminHandler.GetAdminByID)
	adminsProtected.Put("/:id", adminHandler.UpdateAdmin)
	adminsProtected.Delete("/:id", adminHandler.DeleteAdmin)

	// Mentor Routes
	mentors := v1.Group("/mentors")
	mentors.Get("/", mentorHandler.GetAllMentors)
	mentors.Get("/:id", mentorHandler.GetMentorByID)

	mentorsProtected := mentors.Group("/")
	mentorsProtected.Use(middleware.Protected(), middleware.AdminProtected())
	mentorsProtected.Post("/", mentorHandler.CreateMentor)
	mentorsProtected.Put("/:id", mentorHandler.UpdateMentor)
	mentorsProtected.Delete("/:id", mentorHandler.DeleteMentor)

	// Course Routes
	courses := v1.Group("/courses")
	courses.Get("/", courseHandler.GetAllCourse)
	courses.Get("/:id", courseHandler.GetCourseByID)

	coursesProtected := courses.Group("/")
	coursesProtected.Use(middleware.Protected(), middleware.AdminProtected())
	coursesProtected.Post("/", courseHandler.CreateCourse)
	coursesProtected.Post("/upload", courseHandler.UploadFile)
	coursesProtected.Put("/:id", courseHandler.UpdateCourse)
	coursesProtected.Delete("/:id", courseHandler.DeleteCourse)

	// User Course Routes (Enrollment)
	userCourses := v1.Group("/user-courses")
	userCourses.Use(middleware.Protected())
	userCourses.Post("/enroll", userCourseHandler.EnrollCourse)
	userCourses.Get("/my-courses", userCourseHandler.GetUserCourseDashboard)
	userCourses.Get("/dashboard", userCourseHandler.GetUserCourseDashboard)
	userCourses.Post("/payment", transactionHandler.PaymentCourse)

	userCourses.Get("/transactions", transactionHandler.GetTransactionHistory)
	userCourses.Get("/pending", userCourseHandler.GetCoursePending)
	userCourses.Delete("/delete", userCourseHandler.DeleteCourse)

	v1.Post("/callback-notification", transactionHandler.GetNotification)

	// Static File Routes
	staticFile := v1.Group("/static")
	staticFile.Get("/public/uploads/courses/:path", services.StaticFile)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
