package model

import "time"

type SettlementGroup struct {
	ID      uint
	MsgID   string
	CtrlSum float64
	CreDtTm *time.Time
	NbOfTxs uint
}
