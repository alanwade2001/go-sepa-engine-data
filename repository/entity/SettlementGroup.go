package entity

import (
	"time"

	"gorm.io/gorm"
)

type SettlementGroup struct {
	Model       *gorm.Model `gorm:"embedded"`
	MsgID       string
	CtrlSum     float64
	CreDtTm     *time.Time
	NbOfTxs     uint
	ExecutionID uint
	Execution   *Execution
	GrpHdr      string
	DocID       uint
}
