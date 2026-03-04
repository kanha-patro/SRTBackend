package repository

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/akpatri/srt/internal/domain"
)



type auditRepository struct {
	db *gorm.DB
}

func NewAuditRepository(db *gorm.DB) AuditRepository {
	return &auditRepository{db: db}
}

func (r *auditRepository) CreateAuditLog(ctx context.Context, log *domain.AuditLog) error {
	log.CreatedAt = time.Now()
	
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *auditRepository) GetAuditLogs(ctx context.Context, orgID string, limit, offset int) ([]domain.AuditLog, error) {
	var logs []domain.AuditLog
	err := r.db.WithContext(ctx).
		Where("org_id = ?", orgID).
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Find(&logs).Error
	return logs, err
}