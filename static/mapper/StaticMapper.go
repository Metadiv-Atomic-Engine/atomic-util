package mapper

import (
	"io"
	"path/filepath"
	"time"

	"github.com/Metadiv-Atomic-Engine/atomic-util/static/config"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/service"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/gin-gonic/gin"
)

var StaticMapper = new(staticMapper)

type staticMapper struct{}

func (m *staticMapper) FromUpdateRequest(req *request.StaticInfoUpdate, st *entity.Static) *entity.Static {
	st.Filename = req.Filename
	st.Public = req.Public
	return st
}

func (m *staticMapper) FromContext(ctx *gin.Context) (static *entity.Static, content []byte) {
	content, filename, size, err := m.getStaticContentFromRequest(ctx)
	if err != nil {
		return nil, nil
	}
	content = service.EncryptService.Encrypt(content)
	static = &entity.Static{
		Filename:  filename[:len(filename)-len(filepath.Ext(filename))],
		FileType:  filepath.Ext(filename),
		Size:      size,
		Public:    m.getIsPublicFromRequest(ctx),
		Pinned:    m.getIsPinnedFromRequest(ctx),
		ExpiredAt: time.Now().Add(time.Minute * time.Duration(atomic.Engine.EnvInt(config.STATIC_REVERSE_TIME_IN_MINUTES))).Unix(), // the file will be deleted if it is not being pinned
	}
	static.InitUUID()
	return static, content
}

func (s *staticMapper) getStaticContentFromRequest(
	ctx *gin.Context) (content []byte, filename string, size int64, err error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return nil, "", 0, err
	}

	f, err := file.Open()
	if err != nil {
		return nil, "", 0, err
	}
	defer f.Close()

	content, err = io.ReadAll(f)
	if err != nil {
		return nil, "", 0, err
	}

	return content, file.Filename, file.Size, nil
}

func (s *staticMapper) getIsPublicFromRequest(ctx *gin.Context) bool {
	return ctx.PostForm("public") == "true"
}

func (s *staticMapper) getIsPinnedFromRequest(ctx *gin.Context) bool {
	return ctx.PostForm("pinned") == "true"
}
