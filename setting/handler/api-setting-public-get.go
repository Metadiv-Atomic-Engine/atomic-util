package handler

import (
	"strings"

	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/repo"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/service"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/sql"
)

var API_SETTING_PUBLIC_GET = atomic.NewApiHandler(
	module.Setting,
	"CiYnTyjHfgIjEMDBGyOnC",
	"Public workspace settings get",
	func(ctx *atomic.Context[request.SettingPublicGet]) {

		keys := strings.Split(ctx.Request.Keys, ",")

		check := service.SettingService.CheckWorkspaceKeyIsPublic(keys)
		if check != "" {
			ctx.Err(errs.SETTING_KEY_NOT_FOUND, "("+check+")")
			return
		}

		/*
			We ensure that the workspace has settings before we proceed
		*/
		set := repo.SystemSettingRepo.FindOne(ctx.DB, sql.Eq("workspace_id", ctx.Request.Workspace))
		if set == nil {
			ctx.Err(errs.WORKSPACE_HAS_NO_SETTINGS)
			return
		}

		settings := service.SettingService.GetSettings(ctx.Request.Workspace, keys)
		ctx.OK(settings)
	},
	&atomic.TypescriptOpt{
		FunctionName: "getPublicWorkspaceSettings",
		Models:       []any{request.SettingPublicGet{}},
		Forms:        []string{"keys"},
		Response:     "SystemSetting[]",
	},
)
