package models

import (
	"database/sql/driver"
	"encoding/json"
)

type SourceSecurity struct {
	Token  string `json:"token"`
	Crt    string `json:"crt"`
	Key    string `json:"key"`
	Method string `gorm:"type:varchar(32); not null;" json:"security_method"`
}

type Source struct {
	ID            uint64         `gorm:"primaryKey; autoIncrement" json:"id" `
	Name          string         `gorm:"type:varchar(32); not null; uniqueIndex:name_server; priority:2" json:"name"` // binding:"required"
	Server        string         `gorm:"type:varchar(64); not null; uniqueIndex:name_server; priority:1" json:"server"`
	Security      SourceSecurity `gorm:"type:json; not null;" json:"security,omitempty"` // 不要返回的值 使用*引用
	CpuOccupy     string         `json:"cpu_occupy,omitempty"`                           // 当前集群中被使用的值，两部分：1:分配给wk， 2:集群本身自用
	CpuFree       string         `json:"cpu_free,omitempty"`                             // 空闲值
	MemOccupy     string         `json:"mem_occupy,omitempty"`
	MemFree       string         `json:"mem_free,omitempty"`
	StorageOccupy string         `json:"storage_occupy,omitempty"`
	StorageFree   string         `gorm:"default:500" json:"storage_free,omitempty"`
	Rsa           string         `json:"rsa,omitempty"`                                                                                                // optServer Rsa用于用户加解密
	Status        string         `gorm:"type:varchar(64); default:Pending" json:"status"`                                                              // optServer操作的状态 pending/terminating/runningjq
	Workspace     []Workspace    `gorm:"many2many:workspace_source; constraint:OnDelete:CASCADE; constraint:OnUpdate:CASCADE" json:"source,omitempty"` // many 2 many
	Namespace     []Namespace    `json:"namespace,omitempty"`                                                                                          // has One
}

func (Source) TableName() string {
	return "source"
}

func (c SourceSecurity) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *SourceSecurity) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
