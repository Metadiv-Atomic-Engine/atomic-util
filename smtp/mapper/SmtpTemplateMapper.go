package mapper

import (
	"github.com/Metadiv-Atomic-Engine/aes"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/dto"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var SmtpTemplateMapper = smtpTemplateMapper{}

type smtpTemplateMapper struct{}

func (m *smtpTemplateMapper) encrypt(text string) []byte {
	return aes.EncryptTextToBytes(text, atomic.Engine.EnvString(atomic.DB_ENCRYPT_KEY))
}

func (m *smtpTemplateMapper) decrypt(bytes []byte) string {
	return aes.DecryptBytesToString(bytes, atomic.Engine.EnvString(atomic.DB_ENCRYPT_KEY))
}

func (m *smtpTemplateMapper) FromCreateRequest(req *request.SmtpTemplateCreate) *entity.SmtpTemplate {
	return &entity.SmtpTemplate{
		Type:       req.Type,
		Name:       m.encrypt(req.Name),
		SubjectEn:  m.encrypt(req.SubjectEn),
		SubjectZht: m.encrypt(req.SubjectZht),
		SubjectZhs: m.encrypt(req.SubjectZhs),
		ContentEn:  m.encrypt(req.ContentEn),
		ContentZht: m.encrypt(req.ContentZht),
		ContentZhs: m.encrypt(req.ContentZhs),
	}
}

func (m *smtpTemplateMapper) FromUpdateRequest(e *entity.SmtpTemplate, req *request.SmtpTemplateUpdate) *entity.SmtpTemplate {
	e.Type = req.Type
	e.Name = m.encrypt(req.Name)
	e.SubjectEn = m.encrypt(req.SubjectEn)
	e.SubjectZht = m.encrypt(req.SubjectZht)
	e.SubjectZhs = m.encrypt(req.SubjectZhs)
	e.ContentEn = m.encrypt(req.ContentEn)
	e.ContentZht = m.encrypt(req.ContentZht)
	e.ContentZhs = m.encrypt(req.ContentZhs)
	return e
}

func (m *smtpTemplateMapper) ToDTO(e *entity.SmtpTemplate) *dto.SmtpTemplate {
	d := &dto.SmtpTemplate{}
	d.ID = e.ID
	d.Type = e.Type
	d.Name = m.decrypt(e.Name)
	d.SubjectEn = m.decrypt(e.SubjectEn)
	d.SubjectZht = m.decrypt(e.SubjectZht)
	d.SubjectZhs = m.decrypt(e.SubjectZhs)
	d.ContentEn = m.decrypt(e.ContentEn)
	d.ContentZht = m.decrypt(e.ContentZht)
	d.ContentZhs = m.decrypt(e.ContentZhs)
	return d
}

func (m *smtpTemplateMapper) ToDTOs(es []entity.SmtpTemplate) []dto.SmtpTemplate {
	DTOs := make([]dto.SmtpTemplate, len(es))
	for i, e := range es {
		DTOs[i] = *m.ToDTO(&e)
	}
	return DTOs
}
