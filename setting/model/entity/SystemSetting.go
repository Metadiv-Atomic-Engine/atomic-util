package entity

import "strconv"

type SystemSetting struct {
	Key         string `gorm:"primaryKey;not null" json:"key"`
	Type        string `gorm:"not null" json:"type"` // string, int, float, bool, image
	Value       string `json:"value"`                // if the type is image, the value is the image uuid
	WorkspaceID uint   `json:"workspace_id"`         // 0 means global settings
	Public      bool   `json:"public"`
}

func (s *SystemSetting) String() string {
	return s.Value
}

func (s *SystemSetting) Int() int {
	i, err := strconv.Atoi(s.Value)
	if err != nil {
		return 0
	}
	return i
}

func (s *SystemSetting) Float() float64 {
	f, err := strconv.ParseFloat(s.Value, 64)
	if err != nil {
		return 0
	}
	return f
}

func (s *SystemSetting) Bool() bool {
	return s.Value == "true"
}

func (s *SystemSetting) Image() string {
	return s.Value
}
