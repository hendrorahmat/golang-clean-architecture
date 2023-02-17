package services

import domain_errors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"

type ISentEmailService interface {
	Send() domain_errors.DomainError
	GetBcc() []string
	GetSubject() string
	GetCc() []string
	GetMessage() string
	GetAttachments() []string
}

type SendEmailService struct {
	from        string
	to          []string
	cc          []string
	bcc         []string
	subject     string
	messages    string
	attachments []string
}

func (s *SendEmailService) GetSubject() string {
	return s.subject
}

func (s *SendEmailService) GetCc() []string {
	return s.cc
}

func (s *SendEmailService) GetMessage() string {
	return s.messages
}

func (s *SendEmailService) GetAttachments() []string {
	return s.attachments
}

func (s *SendEmailService) Send() domain_errors.DomainError {
	//TODO implement me
	panic("implement me")
}

func (s *SendEmailService) GetBcc() []string {
	return s.bcc
}

func NewSendEmailService(from string,
	to []string,
	cc []string,
	bcc []string,
	subject string,
	messages string,
	attachments []string,
) ISentEmailService {
	return &SendEmailService{from: from,
		to:          to,
		cc:          cc,
		bcc:         bcc,
		subject:     subject,
		messages:    messages,
		attachments: attachments,
	}
}
