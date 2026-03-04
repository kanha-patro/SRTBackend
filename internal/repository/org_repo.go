package repository

import (
	"github.com/akpatri/srt/internal/domain"
	"gorm.io/gorm"
)

type orgRepository struct {
	db *gorm.DB
}

func NewOrgRepository(db *gorm.DB) OrgRepository {
	return &orgRepository{db: db}
}

func (r *orgRepository) Create(org *domain.Org) error {
	return r.db.Create(org).Error
}

func (r *orgRepository) GetByID(id string) (*domain.Org, error) {
	var org domain.Org
	if err := r.db.First(&org, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &org, nil
}

func (r *orgRepository) Update(org *domain.Org) error {
	return r.db.Save(org).Error
}

func (r *orgRepository) Delete(id string) error {
	return r.db.Delete(&domain.Org{}, id).Error
}

func (r *orgRepository) GetAll() ([]domain.Org, error) {
	var orgs []domain.Org
	if err := r.db.Find(&orgs).Error; err != nil {
		return nil, err
	}
	return orgs, nil
}
