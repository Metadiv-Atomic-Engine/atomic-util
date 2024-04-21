package request

import "github.com/Metadiv-Atomic-Engine/atomic/base"

type StaticList struct {
	base.RequestListing
	UUIDs    string `form:"uuids"` // split by comma
	Unpinned bool   `form:"unpinned"`
}
