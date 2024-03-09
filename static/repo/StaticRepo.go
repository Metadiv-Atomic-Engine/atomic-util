package repo

import (
	"time"

	"github.com/Metadiv-Atomic-Engine/atomic-util/static/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
	"github.com/Metadiv-Atomic-Engine/sql"
	"gorm.io/gorm"
)

var StaticRepo = new(staticRepo)

type staticRepo struct {
	base.Repository[entity.Static]
}

func (r *staticRepo) FindByUUID(tx *gorm.DB, uuid string, workspaceID ...uint) *entity.Static {
	return r.FindOne(tx, sql.Eq("uuid", uuid), workspaceID...)
}

func (r *staticRepo) FindPublicByUUID(tx *gorm.DB, uuid string) *entity.Static {
	st, err := sql.FindOne[entity.Static](tx, sql.And(sql.Eq("uuid", uuid), sql.Eq("public", true)))
	if err != nil {
		return nil
	}
	return st
}

func (r *staticRepo) FindAllExpiredNonPinned(tx *gorm.DB) []entity.Static {
	sts, err := sql.FindAll[entity.Static](tx, sql.And(
		sql.Eq("pinned", false),
		sql.Lt("expired_at", time.Now().Unix()),
	))
	if err != nil {
		return nil
	}
	return sts
}
