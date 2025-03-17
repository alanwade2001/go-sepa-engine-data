package repository

import (
	"log/slog"

	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
)

type SettlementGroup struct {
	persist *db.Persist
}

func NewSettlementGroup(persist *db.Persist) *SettlementGroup {
	initiation := &SettlementGroup{
		persist: persist,
	}

	return initiation
}

func (s *SettlementGroup) Perist(entity *entity.SettlementGroup) (*entity.SettlementGroup, error) {
	slog.Debug("settlementGroup", "entity", entity)
	tx := s.persist.DB.Save(entity)
	err := tx.Error
	return entity, err
}
