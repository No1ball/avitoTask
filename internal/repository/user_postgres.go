package repository

import (
	"fmt"
	"github.com/No1ball/avitoTask/internal/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserDB struct {
	db *sqlx.DB
}

func NewUserBD(db *sqlx.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

func (r *UserDB) GetUserBalance(id int) (int, error) {
	var user int
	query := fmt.Sprintf("SELECT main_balance FROM %s WHERE id = $1", userTable)
	err := r.db.Get(&user, query, id)
	return user, err
}

func (r *UserDB) AddCashToUser(userId, cost int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	createUserQuery := fmt.Sprintf("INSERT INTO %s (id, main_balance) values ($1, $2)", userTable)
	_, err = tx.Exec(createUserQuery, userId, cost)
	if err != nil {
		tx.Rollback()
		return err
	}

	todayData := time.Now()
	createRefillQuery := fmt.Sprintf("INSERT INTO %s  (user_id, cost, refill_date) values ($1, $2, $3)", refillTable)
	_, err = tx.Exec(createRefillQuery, userId, cost, todayData)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *UserDB) ReserveCash(userId, orderId, serviceId, cost int, serviceName, description string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	createOrderQuery := fmt.Sprintf("INSERT INTO %s (id, cost) values ($1, $2)", orderTable)
	_, err = tx.Exec(createOrderQuery, orderId, cost)
	if err != nil {
		tx.Rollback()
		return err
	}

	createAccountingQuery := fmt.Sprintf("INSERT INTO %s (id, order_id, service_name, description, service_date) values ($1, $2, $3, $4, $5)", accountingTable)
	_, err = tx.Exec(createAccountingQuery, serviceId, orderId, serviceName, description, time.Now())
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf(
		`UPDATE %s  
				SET reserve_balance = reserve_balance + $1, main_balance = main_balance - $1
				WHERE id = $2`,
		userTable)
	_, err = tx.Exec(query, cost, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *UserDB) GetUser(id int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("select * from %s  where id = $1", userTable)
	row := r.db.QueryRow(query, id)
	err := row.Scan(&user)
	return user, err
}

func (r *UserDB) AddCashToUserWithUpdate(userId, cost int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	todayData := time.Now()
	createRefillQuery := fmt.Sprintf("INSERT INTO %s  (user_id ,cost, refill_date) values ($1, $2, $3)", refillTable)
	_, err = tx.Exec(createRefillQuery, userId, cost, todayData)
	if err != nil {
		tx.Rollback()
		return err
	}
	createUserQuery := fmt.Sprintf("UPDATE %s SET main_balance = main_balance + $1 where id = $2", userTable)
	_, err = tx.Exec(createUserQuery, cost, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
