package models

import (
	"github.com/gogo/protobuf/test"
	"time"
)

type WkSourceQuota struct {
	// 当前Source(optServer) 绑定给workspace
	ID            int64     `json:"id"`
	WorkspaceId   int64     `json:"workspace_id"`
	Workspace     Workspace `gorm:"foreignKey:WorkspaceId; constraint:OnDelete:CASCADE;" json:"workspace"`
	SourceId      int64     `json:"source_id"`
	Source        Source    `gorm:"foreignKey:SourceId; constraint:OnDelete:CASCADE;" json:"source"` // gorm:"constraint:OnDelete:CASCADE"
	CpuFree       string    `json:"cpu_free"`
	CpuOccupy     string    `json:"cpu_occupy"` // 分给wk的值，被占用的
	MemFree       string    `json:"mem_free"`
	MemOccupy     string    `json:"mem_occupy"`
	StorageFree   string    `json:"storage_free"`
	StorageOccupy string    `json:"storage_occupy"`
}

func (WkSourceQuota) TableName() string {
	return "workspaceQuota"
}

type auditLog struct {
	ID       int64     `json:"id"`
	Uuid     test.Uuid `gorm:"index;" json:"uuid"`
	Method   string    `json:"method"`
	User     string    `json:"user"`
	Msg      string    `json:"msg"`
	CreateAt time.Time `json:"create_at"`
	Result   bool      `json:"result"`
}
