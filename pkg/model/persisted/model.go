package persisted

import (
	"gorm.io/gorm"
	"time"
)

type HiddenFieldsModel struct {
	ID        uint           `json:"-" gorm:"primary_key"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" sql:"index"`
}
