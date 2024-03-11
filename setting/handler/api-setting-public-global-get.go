package handler

import (
	"strings"

	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/service"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var API_SETTING_PUBLIC_GLOBAL_GET = atomic.NewApiHandler(
	module.Setting,
	"inP8bfiKnGE6Shl0YxvwR",
	"Public global settings get",
	func(ctx *atomic.Context[request.SettingGet]) {
		keys := strings.Split(ctx.Request.Keys, ",")

		check := service.SettingService.CheckKeyExistsForGlobal(keys)
		if check != "" {
			ctx.Err(errs.SETTING_KEY_NOT_FOUND, "("+check+")")
			return
		}

		settings := service.SettingService.GetSettings(0, keys)
		ctx.OK(settings)
	},
	&atomic.TypescriptOpt{
		FunctionName: "getPublicGlobalSettings",
		Models:       []any{request.SettingGet{}},
		Forms:        []string{"keys"},
		Response:     "SystemSetting[]",
	},
)
