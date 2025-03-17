package entity

import "gorm.io/gorm"

type Execution struct {
	Model *gorm.Model `gorm:"embedded"`
}
