package models

type Namespace struct {
	ID                uint64    `json:"id"`
	Name              string    `gorm:"type:varchar(32); not null;uniqueIndex:name_workspace_source" json:"name"`
	WorkspaceID       uint64    `json:"workspace_id,omitempty"`
	Workspace         Workspace `gorm:"foreignKey:WorkspaceID; constraint:OnDelete:CASCADE; constraint:OnUpdate:CASCADE" json:"workspace"`
	SourceID          uint64    `json:"source_id,omitempty"`
	Source            Source    `gorm:"foreignKey:SourceID; constraint:OnDelete:CASCADE; constraint:OnUpdate:CASCADE" json:"source"`
	NsQuotaCpu        string    `json:"ns_quota_cpu"`
	NsQuotaMem        string    `json:"ns_quota_mem"`
	NsQuotaStorage    string    `json:"ns_quota_storage"`
	NsLimitsCpu       string    `json:"ns_limits_cpu" `
	NsLimitsMem       string    `json:"ns_limits_mem"`
	NsRequestStorage  string    `json:"ns_request_storage" default:"1Ti"`
	NsPvcStorage      string    `json:"ns_pvc_storage" default:"200Gi"`
	PodLimitCpu       string    `json:"pod_limit_cpu" default:"4"`
	PodLimitMem       string    `json:"pod_limit_mem" default:"4Gi"`
	ContainerLimitCpu string    `json:"container_limit_cpu" default:"1"`
	ContainerLimitMem string    `json:"container_limit_mem" default:"200Mi"`
	Account           []Account `gorm:"many2many:account_namespace; constraint:OnDelete:CASCADE; constraint:OnUpdate:CASCADE" json:"account"`
}

func (Namespace) TableName() string {
	return "namespace"
}
