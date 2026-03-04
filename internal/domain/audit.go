package domain

import (
	"time"
)

// Audit represents the structure for audit logs.
type Audit struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Action    string    `json:"action"`
	Actor     string    `json:"actor"`
	TargetID  string    `json:"target_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// AuditLog represents the interface for audit logging.
type AuditLog interface {
	LogAction(action string, actor string, targetID string) error
}

// NewAudit creates a new audit log entry.
func NewAudit(action, actor, targetID string) *Audit {
	return &Audit{
		Action:   action,
		Actor:    actor,
		TargetID: targetID,
		CreatedAt: time.Now(),
	}
}