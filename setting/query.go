package setting

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/sql"
)

/*
Get setting by key and workspaceID
*/
func GetSetting(key string, workspaceID uint) *entity.SystemSetting {
	return repo.SystemSettingRepo.FindOne(atomic.Engine.DB, sql.And(
		sql.Eq("key", key),
		sql.Eq("workspace_id", workspaceID),
	))
}

/*
Get settings by set of keys and workspaceID
*/
func GetSettings(keys []string, workspaceID uint) map[string]entity.SystemSetting {
	settings := repo.SystemSettingRepo.FindAll(atomic.Engine.DB, sql.And(
		sql.In("key", keys),
		sql.Eq("workspace_id", workspaceID),
	), "")
	result := make(map[string]entity.SystemSetting)
	for i := range settings {
		result[settings[i].Key] = settings[i]
	}
	return result
}
