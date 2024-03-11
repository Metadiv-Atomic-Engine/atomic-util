package handler

import (
	"strings"

	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/service"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var API_SETTING_WORKSPACE_GET = atomic.NewApiHandler(
	module.Setting,
	"8fAGa3t0FSf6DytfMJg6V",
	"Workspace user get settings",
	func(ctx *atomic.Context[request.SettingGet]) {

		if ctx.WorkspaceID() == 0 {
			ctx.Forbidden("No workspace")
			return
		}

		keys := strings.Split(ctx.Request.Keys, ",")

		check := service.SettingService.CheckKeyExistsForWorkspace(keys)
		if check != "" {
			ctx.Err(errs.SETTING_KEY_NOT_FOUND, "("+check+")")
			return
		}

		settings := service.SettingService.GetSettings(ctx.WorkspaceID(), keys)
		ctx.OK(settings)
	},
	&atomic.TypescriptOpt{
		FunctionName: "getWorkspaceSettings",
		Models:       []any{entity.SystemSetting{}},
		Forms:        []string{"keys"},
		Response:     "SystemSetting[]",
	},
)
