package entity

import (
	"fmt"

	"gorm.io/gorm"
)

type Settlement struct {
	Model             *gorm.Model `gorm:"embedded"`
	EndToEndID        string
	TxID              string
	CdtTrfTxInf       string
	TransactionID     uint
	Transaction       *Transaction
	SettlementGroupID uint
	SettlementGroup   *SettlementGroup
}

func (t Settlement) String() string {
	return fmt.Sprintf("ID=[%v], EndToEndID=[%v], TxID=[%v]", t.Model.ID, t.EndToEndID, t.TxID)
}
