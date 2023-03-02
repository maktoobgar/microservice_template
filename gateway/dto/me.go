package dto

import "service/auth/models"

type MeResponse struct {
	User models.User `json:"user"`
}
