package request

import "github.com/Metadiv-Atomic-Engine/atomic/base"

type SmtpAccountUpdate struct {
	base.RequestIDPath
	SmtpAccountCreate
}
