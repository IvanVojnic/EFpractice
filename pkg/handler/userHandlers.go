package handler

import (
	"EFpractic2/models"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) createUser(c echo.Context) error {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		log.WithFields(log.Fields{
			"Error Bind json while creating user": err,
			"user":                                user,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	err = h.services.UserAct.CreateUser(c.Request().Context(), user)
	if err != nil {
		log.WithFields(log.Fields{
			"Error create user": err,
			"user":              user,
		}).Info("CREATE USER request")
		return echo.NewHTTPError(http.StatusBadRequest, "user creating failed")
	}
	return c.String(http.StatusOK, "user created")
}

func (h *Handler) getUser(c echo.Context) error {
	userId := c.QueryParam("id")
	var userIdNum int
	userIdNum, _ = strconv.Atoi(userId)
	user, err := h.services.UserAct.GetUser(c.Request().Context(), userIdNum)
	if err != nil {
		log.WithFields(log.Fields{
			"Error get user": err,
			"user":           user,
		}).Info("GET USER request")
		return echo.NewHTTPError(http.StatusBadRequest, "user getting failed")
	}
	fmt.Sprintf("user: %s", user)
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c echo.Context) error {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		log.WithFields(log.Fields{
			"Error Bind json while updating user": err,
			"user":                                user,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	err = h.services.UserAct.UpdateUser(c.Request().Context(), user)
	if err != nil {
		log.WithFields(log.Fields{
			"Error update user": err,
			"user":              user,
		}).Info("UPDATE USER request")
		return echo.NewHTTPError(http.StatusBadRequest, "user updating failed")
	}
	return c.String(http.StatusOK, "user updated")
}

func (h *Handler) deleteUser(c echo.Context) error {
	userId := c.QueryParam("id")
	var userIdNum int
	userIdNum, _ = strconv.Atoi(userId)
	err := h.services.UserAct.DeleteUser(c.Request().Context(), userIdNum)
	if err != nil {
		log.WithFields(log.Fields{
			"Error get user": err,
			"user ID":        userId,
		}).Info("DELETE USER request")
		return echo.NewHTTPError(http.StatusBadRequest, "user deleting failed")
	}
	return c.String(http.StatusOK, "user deleted")
}

func (h *Handler) getAllUsers(c echo.Context) error {
	users, err := h.services.UserAct.GetAllUsers(c.Request().Context())
	if err != nil {
		log.WithFields(log.Fields{
			"Error get all users": err,
			"users":               users,
		}).Info("GET ALL USER request")
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, users)
}
