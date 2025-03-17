package repository

import (
	"log/slog"

	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
)

type Execution struct {
	persist *db.Persist
}

func NewExecution(persist *db.Persist) *Execution {
	execution := &Execution{
		persist: persist,
	}

	return execution
}

func (s *Execution) Perist(entity *entity.Execution) (*entity.Execution, error) {
	slog.Debug("persist", "entity", entity)

	tx := s.persist.DB.Save(entity)
	err := tx.Error

	return entity, err
}
