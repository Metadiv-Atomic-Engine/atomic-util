package module

import "github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/entity"

/*
We store the registered settings in the system memory.
*/
var RegisteredSettingMap = make(map[string]entity.SystemSetting)
