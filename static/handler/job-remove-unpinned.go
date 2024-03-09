package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/repo"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/service"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var JOB_REMOVE_UNPINNED = atomic.NewJobHandler(
	module.Static,
	"CwNy6ujIy2CFmoyd5qcgd",
	"Remove unpinned job",
	func() {
		sts := repo.StaticRepo.FindAllExpiredNonPinned(atomic.Engine.DB)
		tx := atomic.Engine.DB.Begin()
		for i := range sts {
			repo.StaticRepo.Delete(tx, &sts[i])
			service.ContentFileService.Delete(sts[i].WorkspaceID, sts[i].UUID)
		}
		tx.Commit()
	},
)
