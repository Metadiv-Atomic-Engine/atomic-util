package entity

type SystemSetting struct {
	Key         string `gorm:"primaryKey;not null" json:"key"`
	Type        string `gorm:"not null" json:"type"` // string, int, float, bool
	Value       string `json:"value"`
	WorkspaceID uint   `json:"workspace_id"` // 0 means global settings
	Public      bool   `json:"public"`
}
