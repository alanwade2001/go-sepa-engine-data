package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PaymentGroup struct {
	Model   *gorm.Model `gorm:"embedded"`
	InitnID uint
	MsgID   string
	CtrlSum float64
	CreDtTm *time.Time
	NbOfTxs uint
	State   string
	DocID   uint
	GrpHdr  string
}

func (i PaymentGroup) String() string {
	return fmt.Sprintf("ID=[%d], InitnID=[%d], MsgId=[%s], CtrlSum=[%f], NbOfTxs=[%d], State=[%s], DocID=[%d]", i.Model.ID, i.InitnID, i.MsgID, i.CtrlSum, i.NbOfTxs, i.State, i.DocID)
}
