package request

import "github.com/Metadiv-Atomic-Engine/atomic/base"

type StaticInfoUpdate struct {
	base.RequestUUIDPath
	Filename string `json:"filename"`
	Public   bool   `json:"public"`
}
