package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/migrator"
)

// Migrate performs the database migrations.
func Migrate(db *gorm.DB) error {
	// Run migrations
	err := db.AutoMigrate(
		// Add your models here
		&Org{},
		&Route{},
		&Stop{},
		&Trip{},
		&Driver{},
		&Location{},
		&OTP{},
		&Audit{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migrated successfully")
	return nil
}

// Rollback performs the rollback of the last migration.
func Rollback(db *gorm.DB) error {
	// Rollback logic can be implemented here if needed
	return fmt.Errorf("rollback not implemented")
}

// SeedDatabase seeds the database with initial data.
func SeedDatabase(db *gorm.DB) error {
	// Seed initial data if necessary
	return nil
}