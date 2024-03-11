package setting

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/module"
)

/*
Register a new setting key.
The setting key will store to the system memory first.
Then it will be generated to the database on fly if it's not exist.
*/
func Register(
	key string,
	defaultValue string,
	valueType string, // string, int, float, bool
	public bool,
) (settingKey string) {
	if _, ok := module.RegisteredSettingMap[key]; ok {
		panic("Setting key already exist: " + key)
	}

	module.RegisteredSettingMap[key] = entity.SystemSetting{
		Key:         key,
		Type:        valueType,
		Value:       defaultValue,
		WorkspaceID: 1, // 0 means global settings, 1 means workspace settings, however, we will fill the actual workspace id in database
		Public:      public,
	}

	return key
}

/*
RegisterGlobal a new setting key for global settings,
which is editable by the system admin.
*/
func RegisterGlobal(
	key string,
	defaultValue string,
	valueType string, // string, int, float, bool
	public bool,
) (settingKey string) {
	if _, ok := module.RegisteredSettingMap[key]; ok {
		panic("Setting key already exist: " + key)
	}

	module.RegisteredSettingMap[key] = entity.SystemSetting{
		Key:         key,
		Type:        valueType,
		Value:       defaultValue,
		WorkspaceID: 0, // 0 means global settings
		Public:      public,
	}

	return key
}
