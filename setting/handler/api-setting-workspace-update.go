package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/service"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var API_SETTING_WORKSPACE_UPDATE = atomic.NewApiHandler(
	module.Setting,
	"uZzw9iHepsBR21s6JWvkb",
	"Workspace user update settings",
	func(ctx *atomic.Context[request.SettingUpdate]) {

		if ctx.WorkspaceID() == 0 {
			ctx.Forbidden("No workspace")
			return
		}

		keys := make([]string, 0)
		for _, setting := range ctx.Request.Settings {
			keys = append(keys, setting.Key)
		}

		check := service.SettingService.CheckKeyExistsForWorkspace(keys)
		if check != "" {
			ctx.Err(errs.SETTING_KEY_NOT_FOUND, "("+check+")")
			return
		}

		var keyValue = make(map[string]string)
		for i := range ctx.Request.Settings {
			keyValue[ctx.Request.Settings[i].Key] = ctx.Request.Settings[i].Value
		}

		settings := service.SettingService.UpdateSettings(ctx.WorkspaceID(), keys, keyValue)
		ctx.OK(settings)
	},
	&atomic.TypescriptOpt{
		FunctionName: "updateWorkspaceSettings",
		Models:       []any{request.SettingUpdate{}, entity.SystemSetting{}},
		Body:         "SettingUpdate",
		Response:     "SystemSetting[]",
	},
)
