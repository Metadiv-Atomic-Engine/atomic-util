package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/service"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var API_SETTING_GLOBAL_UPDATE = atomic.NewApiHandler(
	module.Setting,
	"Ty4Sxz37SkdedQygKvEKL",
	"Global settings update",
	func(ctx *atomic.Context[request.SettingUpdate]) {

		keys := make([]string, 0)
		for _, setting := range ctx.Request.Settings {
			keys = append(keys, setting.Key)
		}

		check := service.SettingService.CheckKeyExistsForGlobal(keys)
		if check != "" {
			ctx.Err(errs.SETTING_KEY_NOT_FOUND, "("+check+")")
			return
		}

		var keyValue = make(map[string]string)
		for i := range ctx.Request.Settings {
			keyValue[ctx.Request.Settings[i].Key] = ctx.Request.Settings[i].Value
		}

		settings := service.SettingService.UpdateSettings(0, keys, keyValue)
		ctx.OK(settings)
	},
	&atomic.TypescriptOpt{
		FunctionName: "updateGlobalSettings",
		Models:       []any{request.SettingUpdate{}, entity.SystemSetting{}},
		Body:         "SettingUpdate",
		Response:     "SystemSetting[]",
	},
)
