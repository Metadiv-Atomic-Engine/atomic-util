package request

import "github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/entity"

type SettingUpdate struct {
	Settings []entity.SystemSetting `json:"settings"`
}
