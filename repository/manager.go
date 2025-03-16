package repository

import db "github.com/alanwade2001/go-sepa-db"

type Manager struct {
	PaymentGroup *PaymentGroup
	Payment      *Payment
	Transaction  *Transaction
	Settlement   *Settlement
}

func NewManager(p *db.Persist) *Manager {
	manager := &Manager{
		PaymentGroup: NewPaymentGroup(p),
		Payment:      NewPayment(p),
		Transaction:  NewTransaction(p),
		Settlement:   NewSettlement(p),
	}

	return manager
}
