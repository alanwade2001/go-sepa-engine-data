package model

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"
	"gorm.io/gorm"
)

type Payment struct {
	ID          uint
	PmtInfId    string
	CtrlSum     float64
	NbOfTxs     uint
	ReqdExctnDt *time.Time
	DbtrAcc     *Account
	PmtInf      *pain_001_001_03.PaymentInstructionInformation3
}

func NewPayment(pmtInf *pain_001_001_03.PaymentInstructionInformation3) *Payment {

	nbOfTxs, _ := strconv.ParseUint(pmtInf.NbOfTxs, 10, 64)
	reqdExctnDt, _ := time.Parse("2006-01-02", pmtInf.ReqdExctnDt)

	pmt := &Payment{
		PmtInfId:    pmtInf.PmtInfId,
		CtrlSum:     pmtInf.CtrlSum,
		NbOfTxs:     uint(nbOfTxs),
		ReqdExctnDt: &reqdExctnDt,
		DbtrAcc:     &Account{Nm: pmtInf.Dbtr.Nm, Iban: pmtInf.DbtrAcct.Id.IBAN, Bic: pmtInf.DbtrAgt.FinInstnId.BIC},
		PmtInf:      pmtInf,
	}

	return pmt
}

func NewPayments(pmtInves []*pain_001_001_03.PaymentInstructionInformation3) []*Payment {
	pmts := []*Payment{}

	for _, pmtInf := range pmtInves {
		pmt := NewPayment(pmtInf)
		pmts = append(pmts, pmt)
	}

	return pmts
}

func (p Payment) ToEntity() (ent *entity.Payment, err error) {

	if bytes, err := xml.Marshal(p.PmtInf); err != nil {
		return nil, err
	} else {

		ent = &entity.Payment{
			Model: &gorm.Model{
				ID: p.ID,
			},
			PmtInfID:    p.PmtInfId,
			CtrlSum:     p.CtrlSum,
			NbOfTxs:     p.NbOfTxs,
			ReqdExctnDt: p.ReqdExctnDt,
			Nm:          p.DbtrAcc.Nm,
			Iban:        p.DbtrAcc.Iban,
			Bic:         p.DbtrAcc.Bic,
			PmtInf:      string(bytes),
		}

		return ent, nil
	}

}

func (p *Payment) FromEntity(ent *entity.Payment) error {
	p.CtrlSum = ent.CtrlSum
	p.DbtrAcc = &Account{
		Nm:   ent.Nm,
		Iban: ent.Iban,
		Bic:  ent.Bic,
	}
	p.ID = ent.Model.ID
	p.NbOfTxs = ent.NbOfTxs
	p.PmtInfId = ent.PmtInfID
	p.ReqdExctnDt = ent.ReqdExctnDt
	p.PmtInf = &pain_001_001_03.PaymentInstructionInformation3{}

	if err := xml.Unmarshal([]byte(ent.PmtInf), p.PmtInf); err != nil {
		return err
	}

	return nil
}
