package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/repo"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/service"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var API_ASSET_SERVE = atomic.NewApiHandler(
	module.Static,
	"sbBdj3AcsceA1cxNXnmTP",
	"Serve asset",
	func(ctx *atomic.Context[base.RequestUUIDPath]) {
		static := repo.StaticRepo.FindByUUID(ctx.DB, ctx.Request.UUID, ctx.WorkspaceID())
		if static == nil {
			ctx.Err(errs.STATIC_NOT_FOUND)
			return
		}

		content := service.ContentFileService.Get(static.WorkspaceID, static.UUID)
		if content == nil {
			ctx.Err(errs.STATIC_NOT_FOUND)
			return
		}

		ctx.OKFile(content, static.Filename+"."+static.FileType)
	},
	nil,
)
