package model

import (
	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
	"gorm.io/gorm"
)

type Execution struct {
	ID               uint
	SettlementGroups []*SettlementGroup
}

func (p Execution) ToEntity() (ent *entity.Execution, err error) {

	ent = &entity.Execution{
		Model: &gorm.Model{
			ID: p.ID,
		},
	}

	return ent, nil

}

func (p *Execution) FromEntity(ent *entity.Execution, sgps []*entity.SettlementGroup) error {

	p.ID = ent.Model.ID

	p.SettlementGroups = []*SettlementGroup{}

	for _, sgp := range sgps {
		sg := &SettlementGroup{
			ID:      sgp.Model.ID,
			MsgID:   sgp.MsgID,
			CtrlSum: sgp.CtrlSum,
			CreDtTm: sgp.CreDtTm,
			NbOfTxs: sgp.NbOfTxs,
		}

		p.SettlementGroups = append(p.SettlementGroups, sg)
	}

	return nil
}
