package entity

import (
	"fmt"

	"gorm.io/gorm"
)

type Settlement struct {
	Model         *gorm.Model `gorm:"embedded"`
	EndToEndID    string
	TxID          string
	CdtTrfTxInf   string
	TransactionID uint
	Transaction   *Transaction
}

func (t Settlement) String() string {
	return fmt.Sprintf("ID=[%d], EndToEndID=[%s], TxID=[%s]", t.Model.ID, t.EndToEndID, t.TxID)
}
