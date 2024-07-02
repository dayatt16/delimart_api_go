package controllers

import (
	"Delimart/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct{}

// Login handles the login process.
func (controller *AuthController) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	platform := c.Request().Header.Get("X-Platform") // Read from custom header

	// Determine required role based on platform
	var requiredRole int
	switch platform {
	case "web":
		requiredRole = 2
	case "mobile":
		requiredRole = 1
	case "desktop":
		requiredRole = 3
	default:
		log.Printf("Unknown platform: %s", platform)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid platform"})
	}

	log.Printf("Attempting to login user: %s", username)

	// Retrieve user by username
	user, err := models.GetUserByUsername(username)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	if user == nil {
		log.Printf("Invalid login attempt for username: %s", username)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	// Check if the provided password matches the stored password
	if !models.CheckPassword(user.Password, password) {
		log.Printf("Invalid password for username: %s", username)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}
	// Validate role
	if user.KdRole != requiredRole {
		log.Printf("Unauthorized role for user %s: %d", username, user.KdRole)
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Unauthorized role"})
	}
	if user.Live == 1 {
		log.Printf("Account is currently active for user %s", username)
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Account is currently active"})
	}

	// Update live status to 1
	err = models.UpdateUserLiveStatus(user.KdPegawai, 1)
	if err != nil {
		log.Printf("Failed to update live status for user %s: %v", user.KdPegawai, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user live status"})
	}

	log.Printf("User %s successfully logged in and live status updated", username)

	// Prepare response with user's name and role code
	response := map[string]interface{}{
		"kd_pegawai":   user.KdPegawai,
		"nama_pegawai": user.NamaPegawai,
		"kd_role":      user.KdRole,
		"live":         user.Live,
	}

	return c.JSON(http.StatusOK, response)
}

// Logout handles the logout process.
func (controller *AuthController) Logout(c echo.Context) error {
	kdPegawai := c.Param("kd_pegawai")

	// Update live status to 0 indicating user is logged out
	err := models.UpdateUserLiveStatus(kdPegawai, 0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user live status"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Logout successful"})
}
