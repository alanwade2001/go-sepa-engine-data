package repository

import (
	"log/slog"

	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
)

type Payment struct {
	persist *db.Persist
}

func NewPayment(persist *db.Persist) *Payment {
	payment := &Payment{
		persist: persist,
	}

	return payment
}

func (s *Payment) Perist(entity *entity.Payment) (*entity.Payment, error) {
	slog.Debug("payment", "entity", entity)

	tx := s.persist.DB.Save(entity)
	err := tx.Error

	return entity, err
}

func (s *Payment) GetPaymentsByPaymentGroupID(ID uint) ([]*entity.Payment, error) {

	var payments []*entity.Payment
	if err := s.persist.DB.Where(&entity.Payment{PaymentGroupID: ID}).Find(&payments).Error; err != nil {
		return nil, err
	} else {
		return payments, nil
	}
}
