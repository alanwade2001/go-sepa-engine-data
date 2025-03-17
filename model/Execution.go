package model

import (
	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
	"gorm.io/gorm"
)

type Execution struct {
	ID uint
}

func (p Execution) ToEntity() (ent *entity.Execution, err error) {

	ent = &entity.Execution{
		Model: &gorm.Model{
			ID: p.ID,
		},
	}

	return ent, nil

}

func (p *Execution) FromEntity(ent *entity.Execution) error {

	p.ID = ent.Model.ID

	return nil
}
