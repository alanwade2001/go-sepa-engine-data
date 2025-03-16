package repository

import (
	"log"

	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
)

type Transaction struct {
	persist *db.Persist
}

func NewTransaction(persist *db.Persist) *Transaction {
	ct := &Transaction{
		persist: persist,
	}

	return ct
}

func (s *Transaction) Perist(entity *entity.Transaction) (*entity.Transaction, error) {
	log.Printf("entity: [%v]", entity)
	tx := s.persist.DB.Save(entity)
	err := tx.Error
	return entity, err
}

func (s *Transaction) GetTransactionsByPaymentID(ID uint) ([]*entity.Transaction, error) {

	var transactions []*entity.Transaction
	if err := s.persist.DB.Where(&entity.Transaction{PaymentID: ID}).Find(&transactions).Error; err != nil {
		return nil, err
	} else {
		return transactions, nil
	}
}
