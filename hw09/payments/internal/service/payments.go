package service

import "github.com/jackc/pgx/v4"

type PaymentService struct {
	db *pgx.Conn
}

func NewPaymentService(db *pgx.Conn) PaymentService {
	return PaymentService{
		db,
	}
}

func (s *PaymentService) StorePayment(orderId string, amount int) error {
	return nil
}

func (s *PaymentService) DeletePayment(orderId string) error {
	return nil
}
