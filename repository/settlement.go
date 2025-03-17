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
