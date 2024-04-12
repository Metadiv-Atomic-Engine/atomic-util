package repo

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var SmtpJobRepo = new(smtpJobRepo)

type smtpJobRepo struct {
	base.Repository[entity.SmtpJob]
}
