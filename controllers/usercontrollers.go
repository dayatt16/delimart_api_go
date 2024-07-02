package controllers

import (
	"Delimart/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct{}

func (controller *UserController) GetAllUsers(c echo.Context) error {
    users, err := models.GetAllUsers()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, users)
}

func (controller *UserController) GetUser(c echo.Context) error {
    Id := c.Param("id")
    user, err := models.GetUserByKodeUser(Id)
    if err != nil {
        if err.Error() == "User not found" {
            return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, user)
}

func (controller *UserController) SearchUser(c echo.Context) error {
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Query parameter 'query' is required"})
	}

	hasil, err := models.SearchUser(query)
	if err != nil {
		log.Printf("Error searching user: %v\n", err)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "No matching user found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search user"})
	}

	return c.JSON(http.StatusOK, hasil)
}

func (controller *UserController) CreateUser(c echo.Context) error {
    var user models.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    if err := models.CreateUser(&user); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, user)
}

func (controller *UserController) UpdateUser(c echo.Context) error {
    Id := c.Param("id")
    var user models.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    updatedUser, err := models.UpdateUserByKodeUser(Id, &user)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, updatedUser)
}

func (controller *UserController) DeleteUser(c echo.Context) error {
    Id := c.Param("id")
    if err := models.DeleteUserByKodeUser(Id); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "User successfully deleted"})
}
