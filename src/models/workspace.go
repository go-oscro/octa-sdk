package models

type Workspace struct {
	ID            int64       `gorm:"autoIncrement" json:"id"`
	Name          string      `gorm:"type:varchar(32); not null; unique" json:"name"`
	CpuSet        string      `json:"cpu_set" validate:"required"` // 新增时，不可少
	CpuOccupy     string      `json:"cpu_occupy"`
	MemSet        string      `json:"mem_set" validate:"required"`
	MemOccupy     string      `json:"mem_occupy"`
	StorageSet    string      `json:"storage_set" validate:"required"`
	StorageOccupy string      `json:"storage_occupy"`
	Source        []Source    `gorm:"many2many:workspace_source; constraint:OnDelete:SET NULL;constraint:OnUpdate:CASCADE" json:"source,omitempty"` //
	Namespace     []Namespace `json:"namespace,omitempty"`                                                                                          // gorm:"constraint:OnDelete:SET NULL"
	Account       []Account   `gorm:"many2many:account_workspace; constraint:OnDelete:CASCADE;" json:"account"`                                     //constraint:OnDelete:SET NULL
}

func (Workspace) TableName() string {
	return "workspace"
}

type ModifyWorkspaceSource struct {
	// 编辑workspace与Source资源分配时使用的参数
	WkId         int64  `json:"id"`
	AddSource    int64  `json:"add_source"`
	DelSource    int64  `json:"del_source"`                        // 当DelSource ！= 0时，判断ns数据是否为空
	QuotaCpu     string `json:"quota_cpu" validate:"required"`     // 当前workspace和source绑定Quota值,是source分配给wk使用
	QuotaMem     string `json:"quota_mem" validate:"required"`     // required 当add_source时不可为空
	QuotaStorage string `json:"quota_storage" validate:"required"` // required 当del_source时可不传
}
