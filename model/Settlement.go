package model

import (
	"encoding/xml"

	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
	"github.com/alanwade2001/go-sepa-iso/pacs_008_001_02"
)

type Settlement struct {
	ID          uint
	Amount      float64
	EndToEndID  string
	TxID        string
	CdtTrfTxInf *pacs_008_001_02.CreditTransferTransactionInformation11
}

func (s *Settlement) FromEntity(ent *entity.Settlement) error {

	cdtTrfTxInf := &pacs_008_001_02.CreditTransferTransactionInformation11{}
	if err := xml.Unmarshal([]byte(ent.CdtTrfTxInf), cdtTrfTxInf); err != nil {
		return err
	}

	s.Amount = ent.Amount
	s.CdtTrfTxInf = cdtTrfTxInf
	s.EndToEndID = ent.EndToEndID
	s.TxID = ent.TxID

	return nil

}

func FromEntities(ents []*entity.Settlement) ([]*Settlement, error) {

	mdls := []*Settlement{}

	for _, ent := range ents {
		set := &Settlement{}
		set.FromEntity(ent)

		mdls = append(mdls, set)
	}

	return mdls, nil
}
