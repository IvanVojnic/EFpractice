package handler

import (
	"EFpractic2/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createUser(c echo.Context) error {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		log.Printf("faied %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = h.services.UserAct.CreateUser(c.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.String(http.StatusOK, "user created")
}

func (h *Handler) getUser(c echo.Context) error {
	userId := c.QueryParam("id")
	var userIdNum int
	userIdNum, _ = strconv.Atoi(userId)
	user, err := h.services.UserAct.GetUser(c.Request().Context(), userIdNum)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c echo.Context) error {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		log.Printf("faied %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = h.services.UserAct.UpdateUser(c.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.String(http.StatusOK, "user updated")
}

func (h *Handler) deleteUser(c echo.Context) error {
	userId := c.QueryParam("id")
	var userIdNum int
	userIdNum, _ = strconv.Atoi(userId)
	err := h.services.UserAct.DeleteUser(c.Request().Context(), userIdNum)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.String(http.StatusOK, "user deleted")
}

func (h *Handler) getAllUsers(c echo.Context) error {
	err, users := h.services.UserAct.GetAllUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, users)
}
