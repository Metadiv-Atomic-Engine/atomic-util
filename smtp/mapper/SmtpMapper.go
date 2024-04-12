package mapper

import (
	"github.com/Metadiv-Atomic-Engine/aes"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/dto"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var SmtpAccountMapper = new(smtpAccountMapper)

type smtpAccountMapper struct{}

func (m *smtpAccountMapper) encrypt(text string) []byte {
	return aes.EncryptTextToBytes(text, atomic.Engine.EnvString(atomic.DB_ENCRYPT_KEY))
}

func (m *smtpAccountMapper) decrypt(bytes []byte) string {
	return aes.DecryptBytesToString(bytes, atomic.Engine.EnvString(atomic.DB_ENCRYPT_KEY))
}

func (m *smtpAccountMapper) FromCreateRequest(req *request.SmtpAccountCreate) *entity.SmtpAccount {
	return &entity.SmtpAccount{
		Host:     m.encrypt(req.Host),
		Port:     req.Port,
		User:     m.encrypt(req.User),
		Password: m.encrypt(req.Password),
	}
}

func (m *smtpAccountMapper) FromUpdateRequest(e *entity.SmtpAccount, req *request.SmtpAccountUpdate) *entity.SmtpAccount {
	e.Host = m.encrypt(req.Host)
	e.Port = req.Port
	e.User = m.encrypt(req.User)
	e.Password = m.encrypt(req.Password)
	return e
}

func (m *smtpAccountMapper) ToDTO(e *entity.SmtpAccount) *dto.SmtpAccount {
	return &dto.SmtpAccount{
		ID:       e.ID,
		Host:     m.decrypt(e.Host),
		Port:     e.Port,
		User:     m.decrypt(e.User),
		Password: m.decrypt(e.Password),
	}
}

func (m *smtpAccountMapper) ToDTOs(es []entity.SmtpAccount) []dto.SmtpAccount {
	DTOs := make([]dto.SmtpAccount, len(es))
	for i, e := range es {
		DTOs[i] = *m.ToDTO(&e)
	}
	return DTOs
}
