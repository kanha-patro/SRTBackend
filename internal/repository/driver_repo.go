package repository

import (
	"github.com/akpatri/srt/internal/domain"
	"gorm.io/gorm"
)

type DriverRepository interface {
	Create(driver *domain.Driver) error
	Update(driver *domain.Driver) error
	Delete(driverID string) error
	FindByID(driverID string) (*domain.Driver, error)
	FindAll() ([]domain.Driver, error)
}

type driverRepository struct {
	db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) DriverRepository {
	return &driverRepository{db: db}
}

func (r *driverRepository) Create(driver *domain.Driver) error {
	return r.db.Create(driver).Error
}

func (r *driverRepository) Update(driver *domain.Driver) error {
	return r.db.Save(driver).Error
}

func (r *driverRepository) Delete(driverID string) error {
	return r.db.Delete(&domain.Driver{}, driverID).Error
}

func (r *driverRepository) FindByID(driverID string) (*domain.Driver, error) {
	var driver domain.Driver
	err := r.db.First(&driver, "id = ?", driverID).Error
	if err != nil {
		return nil, err
	}
	return &driver, nil
}

func (r *driverRepository) FindAll() ([]domain.Driver, error) {
	var drivers []domain.Driver
	err := r.db.Find(&drivers).Error
	if err != nil {
		return nil, err
	}
	return drivers, nil
}