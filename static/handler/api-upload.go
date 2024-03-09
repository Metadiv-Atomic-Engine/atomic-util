package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/mapper"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/repo"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/service"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var API_UPLOAD = atomic.NewApiHandler(
	module.Static,
	"qfRXio4obIQGwQreYKrCI",
	"Upload file",
	func(ctx *atomic.Context[base.RequestEmpty]) {
		static, file := mapper.StaticMapper.FromContext(ctx.Gin)
		if file == nil || static == nil {
			ctx.InternalServerErr("Failed to get file")
			return
		}

		static = repo.StaticRepo.Save(ctx.DB, static, ctx.WorkspaceID())
		if static == nil {
			ctx.InternalServerErr("Failed to save static info")
			return
		}

		if !service.ContentFileService.Save(static.UUID, file, ctx.WorkspaceID()) {
			ctx.InternalServerErr("Failed to save file")
			return
		}
		ctx.OK(static)
	},
	nil,
)
