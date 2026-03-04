package repository

import (
	"context"
	"time"

	"github.com/akpatri/srt/internal/domain"
	"gorm.io/gorm"
)

type auditRepository struct {
	db *gorm.DB
}

func NewAuditRepository(db *gorm.DB) AuditRepository {
	return &auditRepository{db: db}
}

func (r *auditRepository) CreateAuditLog(ctx context.Context, a *domain.Audit) error {
	a.CreatedAt = time.Now()
	return r.db.WithContext(ctx).Create(a).Error
}

func (r *auditRepository) GetAuditLogs(ctx context.Context, orgID string, limit, offset int) ([]domain.Audit, error) {
	var logs []domain.Audit
	err := r.db.WithContext(ctx).
		Where("org_id = ?", orgID).
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Find(&logs).Error
	return logs, err
}
