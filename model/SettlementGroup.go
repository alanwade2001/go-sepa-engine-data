package model

import (
	"encoding/xml"
	"time"

	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
	"github.com/alanwade2001/go-sepa-iso/pacs_008_001_02"
)

type SettlementGroup struct {
	ID      uint
	MsgID   string
	CtrlSum float64
	CreDtTm *time.Time
	NbOfTxs uint
}

func (sg *SettlementGroup) FromEntity(ent *entity.SettlementGroup) error {

	gh := &pacs_008_001_02.GroupHeader33{}
	if err := xml.Unmarshal([]byte(ent.GrpHdr), gh); err != nil {
		return err
	}

	sg.ID = ent.Model.ID
	sg.CreDtTm = ent.CreDtTm
	sg.CtrlSum = ent.CtrlSum
	sg.MsgID = ent.MsgID
	sg.NbOfTxs = ent.NbOfTxs

	return nil

}
