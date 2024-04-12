package mapper

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/dto"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/request"
)

var SmtpJobMapper = new(smtpJobMapper)

type smtpJobMapper struct{}

func (m *smtpJobMapper) FromCreateRequest(req *request.SmtpJobCreate) *entity.SmtpJob {
	e := new(entity.SmtpJob)
	e.SetTo(req.To)
	e.SetCc(req.Cc)
	e.SetBcc(req.Bcc)
	e.AccountId = req.AccountId
	e.TemplateId = req.TemplateId
	e.SetValue(req.Value)
	e.Locale = req.Locale
	e.Status = e.StatusPending()
	return e
}

func (m *smtpJobMapper) FromUpdateRequest(e *entity.SmtpJob, req *request.SmtpJobUpdate) *entity.SmtpJob {
	e.SetTo(req.To)
	e.SetCc(req.Cc)
	e.SetBcc(req.Bcc)
	e.AccountId = req.AccountId
	e.TemplateId = req.TemplateId
	e.SetValue(req.Value)
	e.Locale = req.Locale
	e.Status = req.Status
	e.TryTimes = req.TryTimes
	return e
}

func (m *smtpJobMapper) ToDTO(e *entity.SmtpJob) *dto.SmtpJob {
	d := new(dto.SmtpJob)
	d.ID = e.ID
	d.To = e.GetTo()
	d.Cc = e.GetCc()
	d.Bcc = e.GetBcc()
	d.AccountId = e.AccountId
	d.TemplateId = e.TemplateId
	d.Value = e.GetValue()
	d.Locale = e.Locale
	d.TryTimes = e.TryTimes
	d.Status = e.Status
	if e.Account != nil {
		d.Account = SmtpAccountMapper.ToDTO(e.Account)
	}
	if e.Template != nil {
		d.Template = SmtpTemplateMapper.ToDTO(e.Template)
	}
	return d
}

func (m *smtpJobMapper) ToDTOs(es []entity.SmtpJob) []dto.SmtpJob {
	ds := make([]dto.SmtpJob, len(es))
	for i, e := range es {
		ds[i] = *m.ToDTO(&e)
	}
	return ds
}
