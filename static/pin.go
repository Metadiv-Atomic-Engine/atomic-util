package static

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/repo"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/service"
	"gorm.io/gorm"
)

/*
Pin is a function that pins a static asset to a workspace.
Skip if the uuid is empty string.
*/
func Pin(tx *gorm.DB, staticUUID string, workspaceID uint) (ok bool) {
	if staticUUID == "" {
		return true
	}
	static := repo.StaticRepo.FindByUUID(tx, staticUUID, workspaceID)
	if static == nil {
		return false
	}
	if static.Pinned {
		return true
	}

	static.Pinned = true
	return repo.StaticRepo.Save(tx, static, workspaceID) != nil
}

/*
Unpin is a function that deletes a static asset from a workspace.
Skip if the uuid is empty string.
*/
func Unpin(tx *gorm.DB, staticUUID string, workspaceID uint) (ok bool) {
	if staticUUID == "" {
		return true
	}
	static := repo.StaticRepo.FindByUUID(tx, staticUUID, workspaceID)
	if static == nil {
		return true
	}
	if repo.StaticRepo.Delete(tx, static) {
		return service.ContentFileService.Delete(workspaceID, staticUUID)
	}
	return true
}

/*
SmartPin is a function that pins a new static asset and deletes the old one.
*/
func SmartPin(tx *gorm.DB, oldUUID string, newUUID string, workspaceID uint) (ok bool) {
	if oldUUID == newUUID {
		return true
	}
	if oldUUID != "" {
		Unpin(tx, oldUUID, workspaceID)
	}
	return Pin(tx, newUUID, workspaceID)
}
