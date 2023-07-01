package di

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"loco/pkg/model/persisted"
)

func newDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	if err = db.AutoMigrate(
		&persisted.Translation{},
		&persisted.TranslationKey{},
		&persisted.Tag{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}
