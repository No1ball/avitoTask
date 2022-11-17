package models

type User struct {
	id             int `json:"id" db:"id"`
	mainBalance    int `json:"mainBalance" db:"main_balance"`
	reserveBalance int `json:"reserveBalance" db:"reserve_balance"`
}

func (m *User) GetBalance() int {
	return m.mainBalance
}
