package main

import "pizza-tracker-go/internal/models"

type Handler struct {
	orders *models.OrderModels
	users  *models.UserModel
}

func NewHandler(dbModel *models.DBModels) *Handler {
	return &Handler{
		orders: &dbModel.Order,
		users:  &dbModel.User,
	}
}
