package services

import (
	"bytes"
	"html/template"
	"os"
	"project-kelas-santai/internal/config"

	"gopkg.in/gomail.v2"
)

type EmailService interface {
	SendEmail(to string, subject string, body string) error
	SendInvoiceFile(to string, filePath string) error
	SendInvoiceTemplate(to string, data InvoiceData, templatePath string) error
}

type InvoiceData struct {
	InvoiceNumber string
	UserName      string
	Date          string
	CourseName    string
	MentorName    string
	Level         string
	Total         string
	GroupLink     string
}

type emailService struct {
	cfg *config.Config
}

func NewEmailService(cfg *config.Config) EmailService {
	return &emailService{
		cfg: cfg,
	}
}

func (s *emailService) SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.cfg.Email.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(s.cfg.Email.Host, s.cfg.Email.Port, s.cfg.Email.User, s.cfg.Email.Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func (s *emailService) SendInvoiceFile(to string, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return s.SendEmail(to, "Invoice - Kelas Santai", string(content))
}

func (s *emailService) SendInvoiceTemplate(to string, data InvoiceData, templatePath string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	return s.SendEmail(to, "Invoice Pembayaran - Kelas Santai", body.String())
}
