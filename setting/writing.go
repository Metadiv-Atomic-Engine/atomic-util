package setting

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

/*
Save setting
*/
func SaveSetting(setting *entity.SystemSetting) *entity.SystemSetting {
	return repo.SystemSettingRepo.Save(atomic.Engine.DB, setting)
}

/*
Save settings
*/
func SaveSettings(settings []entity.SystemSetting) []entity.SystemSetting {
	return repo.SystemSettingRepo.SaveAll(atomic.Engine.DB, settings)
}
