package models

type UsersOrders struct {
	id      int `json:"id" db:"id"`
	userId  int `json:"userId" db:""`
	orderId int `json:"orderId"`
}
