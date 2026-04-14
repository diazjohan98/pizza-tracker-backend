package models

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DBModels struct {
	Order OrderModels
	User  UserModel
	DB    *gorm.DB
}

func InitDB(dataSourceName string) (*DBModels, error) {
	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	err = db.AutoMigrate(&Order{}, &OrderItem{}, &User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	dbModel := &DBModels{
		DB:    db,
		Order: OrderModels{DB: db},
		User:  UserModel{DB: db},
	}
	return dbModel, nil
}
