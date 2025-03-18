package repository

import (
	"testing"
	"time"

	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-engine-data/migrate"
	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	persist := db.NewPersist()
	migrate.Migrate(persist.DB, &entity.PaymentGroup{}, "PaymentGroup")
	migrate.Migrate(persist.DB, &entity.Payment{}, "Payment")
	migrate.Migrate(persist.DB, &entity.Transaction{}, "Transaction")
	migrate.Migrate(persist.DB, &entity.Settlement{}, "Settlement")
	migrate.Migrate(persist.DB, &entity.SettlementGroup{}, "SettlementGroup")
	migrate.Migrate(persist.DB, &entity.Execution{}, "Execution")

	PaymentGroup := NewPaymentGroup(persist)
	payment := NewPayment(persist)
	transaction := NewTransaction(persist)
	settlement := NewSettlement(persist)
	settlementGroup := NewSettlementGroup(persist)
	execution := NewExecution((persist))

	testCases := []struct {
		desc        string
		settlements []*entity.Settlement
	}{
		{
			desc: "1 transaction",
			settlements: []*entity.Settlement{
				{
					Amount:      10.01,
					EndToEndID:  "e2e-1",
					TxID:        "123",
					CdtTrfTxInf: "dummy",
					//Transaction:     t1,
					SettlementGroup: nil,
				},
			},
		},
		{
			desc:        "0 transaction",
			settlements: []*entity.Settlement{},
		},
		{
			desc: "2 transaction",
			settlements: []*entity.Settlement{
				{
					Amount:      10.01,
					EndToEndID:  "e2e-1",
					TxID:        "123",
					CdtTrfTxInf: "dummy",
					//Transaction:     t1,
					SettlementGroup: nil,
				},
				{
					Amount:      11.01,
					EndToEndID:  "e2e-2",
					TxID:        "1233",
					CdtTrfTxInf: "dummy2",
					//Transaction:     t1,
					SettlementGroup: nil,
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			rct, _ := time.Parse("2006-03-17", "2025-03-17")

			pg1 := &entity.PaymentGroup{
				MsgID:   "msgId-1",
				CtrlSum: 10.01,
				NbOfTxs: 1,
				InitnID: 1,
				State:   "accepted",
				DocID:   1,
				GrpHdr:  "gh",
			}

			_, err := PaymentGroup.Perist(pg1)

			p1 := &entity.Payment{
				PmtInfID:     "pmtInfId-1",
				CtrlSum:      10.01,
				NbOfTxs:      1,
				ReqdExctnDt:  &rct,
				Nm:           "b-1",
				Iban:         "iban-1",
				Bic:          "bic-1",
				PmtInf:       "pmt",
				PaymentGroup: pg1,
			}

			_, err = payment.Perist(p1)

			t1 := &entity.Transaction{
				EndToEndID:  "e2e-1",
				Amt:         10.01,
				Nm:          "A",
				Iban:        "iban-1",
				Bic:         "bic-1",
				CdtTrfTxInf: "tx",
				Payment:     p1,
			}

			_, err = transaction.Perist(t1)

			settlementsSum := 0.0
			for _, s := range tC.settlements {
				s.Transaction = t1
				_, err = settlement.Perist(s)
				assert.NoError(t, err)
				settlementsSum += s.Amount
			}

			e1 := &entity.Execution{}
			_, err = execution.Perist(e1)
			assert.NoError(t, err)

			sg1 := &entity.SettlementGroup{
				MsgID:     "msg-1",
				CtrlSum:   10.1,
				NbOfTxs:   1,
				Execution: e1,
				GrpHdr:    "gh",
				DocID:     4,
			}

			_, err = settlementGroup.Perist(sg1)
			assert.NoError(t, err)

			var rows int64
			rows, err = settlement.UpdateSettlementGroup(sg1)

			nbOfSettlements := int64(len(tC.settlements))
			assert.Equal(t, nbOfSettlements, rows)

			var sum float64
			sum, err = settlement.SumSettlementAmountBySettlementGroupID(sg1.Model.ID)
			assert.NoError(t, err)
			assert.NotNil(t, sum)
			assert.Equal(t, settlementsSum, sum)

		})
	}
}
