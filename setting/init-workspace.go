package setting

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/sql"
)

/*
InitSettingsForWorkspace initializes settings for a workspace.
You may call it after creating a new workspace.
*/
func InitSettingsForWorkspace(workspaceID uint) (ok bool) {

	keys := make([]string, 0)
	for key := range module.RegisteredSettingMap {
		if module.RegisteredSettingMap[key].WorkspaceID == 1 {
			keys = append(keys, key)
		}
	}

	settings := repo.SystemSettingRepo.FindAll(
		atomic.Engine.DB,
		sql.And(
			sql.Eq("workspace_id", workspaceID),
			sql.In("key", keys),
		),
		"",
	)

	existingMap := make(map[string]bool)
	for _, setting := range settings {
		existingMap[setting.Key] = true
	}

	toCreate := make([]entity.SystemSetting, 0)
	for _, key := range keys {
		if _, ok := existingMap[key]; !ok {
			toCreate = append(toCreate, entity.SystemSetting{
				Key:         key,
				Type:        module.RegisteredSettingMap[key].Type,
				Value:       module.RegisteredSettingMap[key].Value,
				WorkspaceID: workspaceID,
				Public:      module.RegisteredSettingMap[key].Public,
			})
		}
	}

	if len(toCreate) > 0 {
		created := repo.SystemSettingRepo.SaveAll(atomic.Engine.DB, toCreate)
		if created == nil {
			return false
		}
	}

	return true
}
