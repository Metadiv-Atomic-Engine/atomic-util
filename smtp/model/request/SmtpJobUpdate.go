package request

import "github.com/Metadiv-Atomic-Engine/atomic/base"

type SmtpJobUpdate struct {
	base.RequestIDPath
	SmtpJobCreate
	TryTimes int    `json:"try_times"`
	Status   string `json:"status"`
}
