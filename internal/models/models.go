package models

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DBModels struct {
	Order OrderModels
}

func InitDB(dataSourceName string) (*DBModels, error) {
	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %v", err)
	}

	err = db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("Failed to migrate database %v", err)
	}

	dbModel := &DBModels{
		Order: OrderModels{DB: db},
	}
	return dbModel, nil
}
