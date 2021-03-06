// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"learn/04_goProject/homework/internal/config"
	"learn/04_goProject/homework/internal/db"
)

// Injectors from wire.go:

func InitApp() (*App, error) {
	configConfig, err := config.New()
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.New(configConfig)
	if err != nil {
		return nil, err
	}
	app := NewApp(sqlDB)
	return app, nil
}
