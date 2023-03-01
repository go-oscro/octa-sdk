package models

type Account struct {
	ID        uint64      `gorm:"primary_key" json:"id" `
	Name      string      `gorm:"type:varchar(64); not null" json:"name"`
	Workspace []Workspace `gorm:"many2many:account_workspace; constraint:OnDelete:CASCADE;" json:"workspace"` // constraint:OnDelete:SET NULL
	Namespace []Namespace `gorm:"many2many:account_namespace; constraint:OnDelete:CASCADE;" json:"namespace"` // constraint:OnDelete:SET NULL

}

func (Account) TableName() string {
	return "account"
}
