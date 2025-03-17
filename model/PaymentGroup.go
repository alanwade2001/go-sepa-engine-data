package model

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"
	"gorm.io/gorm"
)

type PaymentGroup struct {
	ID      uint
	InitnID uint
	MsgID   string
	CreDtTm *time.Time
	CtrlSum float64
	NbOfTxs uint
	State   string
	DocID   uint
	GrpHdr  *pain_001_001_03.GroupHeader32
}

func (i PaymentGroup) String() string {
	return fmt.Sprintf("ID=[%d], InitnID=[%d], MsgId=[%s], CtrlSum=[%f], NbOfTxs=[%d], State=[%s], DocID=[%d]", i.ID, i.InitnID, i.MsgID, i.CtrlSum, i.NbOfTxs, i.State, i.DocID)
}

func NewPaymentGroup(gh *pain_001_001_03.GroupHeader32) *PaymentGroup {

	nbOfTxs, _ := strconv.ParseUint(gh.NbOfTxs, 10, 64)
	creDtTm, _ := time.Parse("2006-01-02T12:00:00", gh.CreDtTm)

	pg := &PaymentGroup{
		MsgID:   gh.MsgId,
		CtrlSum: gh.CtrlSum,
		NbOfTxs: uint(nbOfTxs),
		CreDtTm: &creDtTm,
		GrpHdr:  gh,
	}

	return pg
}

func (p PaymentGroup) ToEntity() (ent *entity.PaymentGroup, err error) {

	if bytes, err := xml.Marshal(p.GrpHdr); err != nil {
		return nil, err
	} else {

		ent = &entity.PaymentGroup{
			Model: &gorm.Model{
				ID: p.ID,
			},
			InitnID: p.InitnID,
			CtrlSum: p.CtrlSum,
			NbOfTxs: p.NbOfTxs,
			CreDtTm: p.CreDtTm,
			MsgID:   p.MsgID,
			State:   p.State,
			DocID:   p.DocID,
			GrpHdr:  string(bytes),
		}

		return ent, nil
	}

}
