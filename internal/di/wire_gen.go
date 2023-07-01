// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"gorm.io/gorm"
)

// Injectors from wire.go:

func GetDB() (*gorm.DB, error) {
	db, err := newDB()
	if err != nil {
		return nil, err
	}
	return db, nil
}
