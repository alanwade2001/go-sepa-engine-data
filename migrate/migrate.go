package migrate

import (
	"log"

	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-engine-data/repository/entity"
	"gorm.io/gorm"

	utils "github.com/alanwade2001/go-sepa-utils"
)

func main() {
	persist := db.NewPersist()
	schemaName := utils.Getenv("DB_SCHEMA", "public")
	if err := persist.DB.Exec("CREATE SCHEMA IF NOT EXISTS " + schemaName).Error; err != nil {
		log.Fatalf("cannot create schema: [%s], error:[%v]", schemaName, err)
		return
	}

	//persist.DB.AutoMigrate(&entity.PaymentGroup{})
	//log.Printf("created table: [%s]", "PaymentGroup")

	//persist.DB.AutoMigrate(&entity.PaymentGroup{})
	//log.Printf("created table: [%s]", "PaymentGroup")
	Migrate(persist.DB, &entity.PaymentGroup{}, "PaymentGroups")
	Migrate(persist.DB, &entity.Payment{}, "Payments")
	Migrate(persist.DB, &entity.Transaction{}, "Transactions")
	Migrate(persist.DB, &entity.Settlement{}, "Settlement")
	Migrate(persist.DB, &entity.SettlementGroup{}, "SettlementGroup")
	Migrate(persist.DB, &entity.Execution{}, "Execution")

}

func Migrate(db *gorm.DB, i interface{}, tableName string) {
	if err := db.AutoMigrate(i); err != nil {
		log.Fatalf("cannot create table: [%s], error:[%v]", tableName, err)
	} else {
		log.Printf("created table: [%s]", tableName)
	}
}
