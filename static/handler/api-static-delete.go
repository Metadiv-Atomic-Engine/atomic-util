package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/repo"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/service"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var API_STATIC_DELETE = atomic.NewApiHandler(
	module.Static,
	"d3scEEu2lBW45iQkXoQ7j",
	"Delete static (including the asset)",
	func(ctx *atomic.Context[base.RequestUUIDPath]) {
		static := repo.StaticRepo.FindByUUID(ctx.DB, ctx.Request.UUID, ctx.WorkspaceID())
		if static == nil {
			ctx.Err(errs.STATIC_NOT_FOUND)
			return
		}

		service.ContentFileService.Delete(static.WorkspaceID, static.UUID)

		if !repo.StaticRepo.Delete(ctx.DB, static) {
			ctx.InternalServerErr()
			return
		}

		ctx.OK(nil)
	},
	&atomic.TypescriptOpt{
		FunctionName: "deleteStatic",
		Paths:        []string{"uuid"},
	},
)
