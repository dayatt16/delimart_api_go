package controllers

import (
    "net/http"
    "Delimart/models"
    "github.com/labstack/echo/v4"
)

type RoleController struct{}

func (controller *RoleController) GetAllRoles(c echo.Context) error {
    roles, err := models.GetAllRoles()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, roles)
}

func (controller *RoleController) GetRoleByKodeRole(c echo.Context) error {
    kdRole := c.Param("kd_role")
    role, err := models.GetRoleByKodeRole(kdRole)
    if err != nil {
        if err.Error() == "Role not found" {
            return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, role)
}

func (controller *RoleController) CreateRole(c echo.Context) error {
    var role models.Role
    if err := c.Bind(&role); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    if err := models.CreateRole(&role); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, role)
}

func (controller *RoleController) UpdateRole(c echo.Context) error {
    kdRole := c.Param("kd_role")
    var role models.Role
    if err := c.Bind(&role); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    updatedRole, err := models.UpdateRoleByKodeRole(kdRole, &role)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, updatedRole)
}

func (controller *RoleController) DeleteRole(c echo.Context) error {
    kdRole := c.Param("kd_role")
    if err := models.DeleteRoleByKodeRole(kdRole); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "Role successfully deleted"})
}
