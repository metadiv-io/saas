package history

import (
	"time"

	"github.com/google/uuid"
	"github.com/metadiv-io/saas/types"
	"gorm.io/gorm"
)

type IHistory interface {
	SetHistory(claims *types.Jwt, action string)
}

type History struct {
	gorm.Model
	TraceID         string `gorm:"not null"` // uuid
	Action          string `gorm:"not null"` // create, update, delete
	EditorUUID      string `gorm:"not null"`
	EditorName      string `gorm:"not null"`
	EditorType      string `gorm:"not null"`
	EditorIP        string `gorm:"not null"`
	EditorUserAgent string `gorm:"not null"`
}

func (h *History) SetHistory(claims *types.Jwt, action string) {
	h.ID = 0
	h.CreatedAt = time.Now()
	h.UpdatedAt = time.Now()
	h.TraceID = uuid.NewString()
	h.Action = action
	if claims != nil {
		h.EditorUUID = claims.UserUUID
		h.EditorName = claims.Username
		h.EditorType = claims.Type
		if len(claims.IPs) > 0 {
			h.EditorIP = claims.IPs[0]
		}
		h.EditorUserAgent = claims.UserAgent
	} else {
		h.EditorUUID = "system"
		h.EditorName = "system"
		h.EditorType = "system"
		h.EditorUserAgent = "system"
	}
}

type HistoryDTO struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	TraceID         string    `json:"trace_id"`
	Action          string    `json:"action"`
	EditorUUID      string    `json:"editor_uuid"`
	EditorName      string    `json:"editor_name"`
	EditorType      string    `json:"editor_type"`
	EditorIP        string    `json:"editor_ip"`
	EditorUserAgent string    `json:"editor_user_agent"`
}
