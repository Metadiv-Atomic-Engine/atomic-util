package service

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/sql"
)

var SettingService = new(settingService)

type settingService struct{}

/*
Check key exists for workspace settings.
return a key that does not exist, or empty string if all keys exist.
*/
func (s *settingService) CheckKeyExistsForWorkspace(keys []string) string {
	for _, key := range keys {
		set, ok := module.RegisteredSettingMap[key]
		if !ok {
			return key
		}
		if set.WorkspaceID == 0 {
			return key
		}
	}
	return ""
}

/*
Check key exists for global settings.
return a key that does not exist, or empty string if all keys exist.
*/
func (s *settingService) CheckKeyExistsForGlobal(keys []string) string {
	for _, key := range keys {
		set, ok := module.RegisteredSettingMap[key]
		if !ok {
			return key
		}
		if set.WorkspaceID != 0 {
			return key
		}
	}
	return ""
}

/*
Check key is public.
return a key that is not public, or empty string if all keys are public.
*/
func (s *settingService) CheckWorkspaceKeyIsPublic(keys []string) string {
	for _, key := range keys {
		set, ok := module.RegisteredSettingMap[key]
		if !ok {
			return key
		}
		if set.WorkspaceID != 0 {
			return key
		}
		if !set.Public {
			return key
		}
	}
	return ""
}

/*
Check global key is public.
return a key that is not public, or empty string if all keys are public.
*/
func (s *settingService) CheckGlobalKeyIsPublic(keys []string) string {
	for _, key := range keys {
		set, ok := module.RegisteredSettingMap[key]
		if !ok {
			return key
		}
		if set.WorkspaceID == 0 {
			return key
		}
		if !set.Public {
			return key
		}
	}
	return ""
}

func (s *settingService) GetSettings(workspaceID uint, keys []string) []entity.SystemSetting {
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
		if created != nil {
			settings = append(settings, created...)
		}
	}

	return settings
}

func (s *settingService) UpdateSettings(workspaceID uint, keys []string, keyValue map[string]string) []entity.SystemSetting {
	settings := repo.SystemSettingRepo.FindAll(
		atomic.Engine.DB,
		sql.And(
			sql.Eq("workspace_id", workspaceID),
			sql.In("key", keys),
		),
		"",
	)

	for i := range settings {
		settings[i].Value = keyValue[settings[i].Key]
	}

	return repo.SystemSettingRepo.SaveAll(atomic.Engine.DB, settings)
}
