package repo

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var SystemSettingRepo = new(systemSettingRepo)

type systemSettingRepo struct {
	base.Repository[entity.SystemSetting]
}
