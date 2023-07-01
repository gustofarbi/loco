//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	wire.Build(newDB)
	return nil, nil
}
