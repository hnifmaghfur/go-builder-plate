package services

import (
	"go-builder-plate/src/auth"
	"go-builder-plate/src/models"
	"go-builder-plate/src/utils"
	"net/http"
)

func Login(payload models.LoginRequest) (int, interface{}) {
	// validate login data
	if payload.Username != "hanif" || payload.Password != "123123" {
		return utils.ResponseError(http.StatusBadRequest, "Invalid request")
	}

	// generate token
	token, err := auth.JWT(payload.Username)
	if err != nil {
		return utils.ResponseError(http.StatusInternalServerError, "Failed to generate token")
	}

	data := map[string]interface{}{
		"token": token,
		"user":  payload.Username,
	}

	return utils.Response(http.StatusOK, "Login successful", data)
}
