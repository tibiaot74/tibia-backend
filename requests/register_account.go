package requests

import "tibia-backend/models"

type RegisterAccountRequest struct {
	Name     *int   `json:"name" binding:"required,gte=100000"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterAccountResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RegisterPlayerRequest struct {
	Name   string      `json:"name" binding:"required"`
	Sex    *models.Sex `json:"sex" binding:"required"`
	Outfit int         `json:"outfit" binding:"required"`
}

type RegisterPlayerResponse struct {
	Id     int        `json:"id"`
	Name   string     `json:"name" binding:"required"`
	Sex    models.Sex `json:"sex" binding:"required"`
	Outfit int        `json:"outfit" binding:"required"`
}
