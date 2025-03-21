package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	Model          *gorm.Model `gorm:"embedded"`
	PmtInfID       string
	CtrlSum        float64
	NbOfTxs        uint
	ReqdExctnDt    *time.Time
	Nm             string
	Iban           string
	Bic            string
	PmtInf         string
	PaymentGroupID uint
	PaymentGroup   *PaymentGroup
}

func (p Payment) String() string {
	return fmt.Sprintf("ID=[%d], PmtInfID=[%s], CtrlSum=[%f], NbOfTxs=[%d], Nm=[%s], Iban=[%s], bic=[%s]", p.Model.ID, p.PmtInfID, p.CtrlSum, p.NbOfTxs, p.Nm, p.Iban, p.Bic)
}
