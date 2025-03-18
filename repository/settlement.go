package repository

import (
	"log/slog"

	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
)

type Settlement struct {
	persist *db.Persist
}

func NewSettlement(persist *db.Persist) *Settlement {
	ct := &Settlement{
		persist: persist,
	}

	return ct
}

func (s *Settlement) Perist(entity *entity.Settlement) (*entity.Settlement, error) {
	slog.Debug("settlement", "entity", entity)
	tx := s.persist.DB.Save(entity)
	err := tx.Error
	return entity, err
}

func (s *Settlement) UpdateSettlementGroup(sg *entity.SettlementGroup) (int64, error) {
	tx := s.persist.DB.Model(Settlement{}).Where("settlement_group_id IS NULL").UpdateColumn("settlement_group_id", sg.Model.ID)

	return tx.RowsAffected, tx.Error
}

func (s *Settlement) FindSettlementsBySettlementGroupID(ID uint) ([]*entity.Settlement, error) {
	var settlements []*entity.Settlement
	if err := s.persist.DB.Where(&entity.Settlement{SettlementGroupID: ID}).Find(&settlements).Error; err != nil {
		return nil, err
	} else {
		return settlements, nil
	}
}
