package entity

import "github.com/Metadiv-Atomic-Engine/atomic/base"

type Static struct {
	base.Model
	base.ModelUUID
	base.ModelWorkspace

	Filename  string `gorm:"not null;" json:"filename"`
	FileType  string `gorm:"not null;" json:"file_type"`
	Size      int64  `gorm:"not null;" json:"size"`
	Public    bool   `gorm:"not null;" json:"public"`
	ExpiredAt int64  `gorm:"not null;" json:"expired_at"`
	Pinned    bool   `gorm:"not null;" json:"pinned"`
}
