package repo

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var SmtpAccountRepo = new(smtpAccountRepo)

type smtpAccountRepo struct {
	base.Repository[entity.SmtpAccount]
}
