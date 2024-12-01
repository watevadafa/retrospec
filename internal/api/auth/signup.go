package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/watevadafa/retrospec/internal/models"
	"github.com/watevadafa/retrospec/internal/models/repositories"
	"github.com/watevadafa/retrospec/internal/types"
	"github.com/watevadafa/retrospec/internal/utils"
)

type SignupRequest struct {
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Plan      string `json:"plan" form:"plan"`
}

var userRepo *repositories.UserRepository

func validateSignupRequest(request SignupRequest) []types.Error {
	var errors []types.Error

	// Required fields validation
	if utils.IsNullOrEmpty(&request.Email) {
		errors = append(errors, types.Error{
			Field:   "email",
			Message: "Email is required",
			Code:    fiber.StatusBadRequest,
		})
	}

	if utils.IsNullOrEmpty(&request.Password) {
		errors = append(errors, types.Error{
			Field:   "password",
			Message: "Password is required",
			Code:    fiber.StatusBadRequest,
		})
	}

	if utils.IsNullOrEmpty(&request.FirstName) {
		errors = append(errors, types.Error{
			Field:   "first_name",
			Message: "First name is required",
			Code:    fiber.StatusBadRequest,
		})
	}

	// If required fields not present raise error
	if len(errors) > 0 {
		return errors
	}
	log.Printf("Required fields validation Successful")

	// Email format validation
	if !utils.IsValidEmail(&request.Email) {
		errors = append(errors, types.Error{
			Field:   "email",
			Message: "Invalid email format",
			Code:    fiber.StatusBadRequest,
		})
	}
	log.Printf("Email format validation Successful")

	// Check if email exists
	existingUser, err := repositories.UserRepo.GetByEmail(request.Email)
	if err != nil {
		log.Fatal(err)
		errors = append(errors, types.Error{
			Field:   "email",
			Message: "Error checking email availability",
			Code:    fiber.StatusInternalServerError,
		})
	}

	if existingUser != nil {
		errors = append(errors, types.Error{
			Field:   "email",
			Message: "Email already registered",
			Code:    fiber.StatusConflict,
		})
	}

	// Password format validation
	personalInfo := types.PersonalInfo{
		FirstName: request.FirstName,
		LastName:  &request.LastName,
		Email:     request.Email,
	}

	if !utils.IsValidPassword(&request.Password, personalInfo) {
		errors = append(errors, types.Error{
			Field:   "password",
			Message: "Password must be between 8 and 64 characters, must contain at least one uppercase letter, one lowercase letter, one number, and one special character. Password cannot contain parts of your name or email.",
			Code:    fiber.StatusBadRequest,
		})
	}
	log.Printf("Email format validation Successful")

	return errors
}

func SignUp(c *fiber.Ctx) error {
	var request SignupRequest
	if err := utils.ParseRequestBody(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(types.Error{
			Field:   "body",
			Message: "Invalid request data",
			Code:    fiber.StatusBadRequest,
		})
	}

	log.Printf("Validating Signup Request")
	if errors := validateSignupRequest(request); len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}
	log.Printf("Successfully Validated Signup Request")

	// Hash password
	hashedPassword, err := utils.HashPassword(&request.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.Error{
			Field:   "password",
			Message: "Error processing password",
			Code:    fiber.StatusInternalServerError,
		})
	}
	log.Printf("Password hashing Successful")

	// Create user
	user := models.NewUser(
		request.Email,
		request.FirstName,
		&request.LastName,
		request.Plan,
	)
	user.PasswordHash = hashedPassword

	// Save user to database
	log.Printf("Saving user to database")
	err = repositories.UserRepo.Create(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.Error{
			Message: "Error saving user to database",
			Code:    fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}
