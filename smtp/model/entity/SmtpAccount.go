package entity

import "github.com/Metadiv-Atomic-Engine/atomic/base"

type SmtpAccount struct {
	base.Model
	base.ModelWorkspace

	Host     []byte `json:"host"`
	Port     int    `json:"port"`
	User     []byte `json:"user"`
	Password []byte `json:"password"`
}
