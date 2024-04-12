package repo

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var SmtpTemplateRepo = new(smtpTemplateRepo)

type smtpTemplateRepo struct {
	base.Repository[entity.SmtpTemplate]
}
