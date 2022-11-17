package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type AccountingDB struct {
	db *sqlx.DB
}

func NewAccountingDB(db *sqlx.DB) *AccountingDB {
	return &AccountingDB{
		db: db,
	}
}

func (r *AccountingDB) RevenueConfirmation(userId, orderId, serviceId, cost int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	updateUserQuery := fmt.Sprintf(
		`UPDATE %s  
				SET reserve_balance = reserve_balance - $1
				WHERE id = $2`,
		userTable)
	_, err = tx.Exec(updateUserQuery, cost, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	createUsersOrdersQuery := fmt.Sprintf("INSERT INTO %s (user_id, order_id) values ($1, $2)", usersOrdersTable)
	_, err = tx.Exec(createUsersOrdersQuery, userId, orderId)
	if err != nil {
		tx.Rollback()
		return err
	}
	todayData := time.Now()
	updateAccountingQuery := fmt.Sprintf(
		`UPDATE %s
				SET service_date = $1, is_completed = TRUE
				WHERE id = $2`, accountingTable)
	_, err = tx.Exec(updateAccountingQuery, todayData, serviceId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
