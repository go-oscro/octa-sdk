package models

import "time"

type AuditLog struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	Uuid        string    `json:"uuid" gorm:"index"`
	Server      string    `json:"server"`
	Method      string    `json:"method"`
	Url         string    `json:"url"`
	OperateKind string    `json:"operateKind" gorm:"index"` // 操作类型服务日志
	Message     string    `json:"message"`
	Result      string    `json:"result"`
	Workspace   string    `json:"workspace" gorm:"index"`
	Namespace   string    `json:"namespace" gorm:"index"`
	CreatedAt   time.Time `json:"create_at"`
	CreateBy    string    `json:"create_by"`
}

func (AuditLog) TableName() string {
	return "auditLog"
}
