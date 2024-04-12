package request

import "github.com/Metadiv-Atomic-Engine/atomic/base"

type SmtpTemplateUpdate struct {
	base.RequestIDPath
	SmtpTemplateCreate
}
