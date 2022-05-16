package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

type PaymentService struct {
	db *pgx.Conn
}

func NewPaymentService(db *pgx.Conn) PaymentService {
	return PaymentService{
		db,
	}
}

func (s *PaymentService) StorePayment(orderId string, amount int) error {
	sqlStmt := `INSERT INTO t_payments (order_id, amount) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := s.db.Exec(context.Background(), sqlStmt, orderId, amount)

	if err != nil {
		return errors.New("could not store payment")
	}

	return nil
}

func (s *PaymentService) DeletePayment(orderId string) error {
	sqlStmt := `DELETE FROM t_payments WHERE order_id = $1`
	_, err := s.db.Exec(context.Background(), sqlStmt, orderId)

	if err != nil {
		return errors.New("could not delete payment")
	}

	return nil
}
