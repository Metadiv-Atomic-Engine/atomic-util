package static

import (
	"path/filepath"
	"time"

	"github.com/Metadiv-Atomic-Engine/atomic-util/static/config"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/repo"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/service"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

/*
This upload is called by another services
*/
func Upload(content []byte, filename string, workspaceID uint, public bool, pinned bool) *entity.Static {
	static := &entity.Static{
		Filename:  filename[:len(filename)-len(filepath.Ext(filename))],
		FileType:  filepath.Ext(filename),
		Size:      int64(len(content)),
		Public:    public,
		Pinned:    pinned,
		ExpiredAt: time.Now().Add(time.Minute * time.Duration(atomic.Engine.EnvInt(config.STATIC_REVERSE_TIME_IN_MINUTES))).Unix(), // the file will be deleted if it is not being pinned
	}
	static.InitUUID()

	if !service.ContentFileService.Save(static.UUID, content, workspaceID) {
		return nil
	}

	static = repo.StaticRepo.Save(atomic.Engine.DB, static, workspaceID)
	if static == nil {
		return nil
	}

	return static
}
