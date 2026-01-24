package services

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"project-kelas-santai/internal/config"

	"gopkg.in/gomail.v2"
)

type EmailService interface {
	SendEmail(to string, subject string, body string, attachments ...string) error
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

func (s *emailService) SendEmail(to string, subject string, body string, attachments ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "kelasantai.bootcamp@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	for _, attachment := range attachments {
		m.Attach(attachment)
	}

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

	// Generate PDF
	pdfFile, err := GenerateInvoicePDF(data)
	if err != nil {
		fmt.Printf("Failed to generate PDF: %v\n", err)
		// Fallback to sending without attachment? Or return error?
		// User asked for file, so let's log and continue, or just return error.
		// For robustness, I'll return the error.
		return err
	}
	defer os.Remove(pdfFile) // Clean up

	return s.SendEmail(to, "Invoice Pembayaran - Kelas Santai", body.String(), pdfFile)
}
func SendEmail(to string, subject string, body string, attachments ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "kelasantai.bootcamp@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	for _, attachment := range attachments {
		m.Attach(attachment)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	d := gomail.NewDialer(cfg.Email.Host, cfg.Email.Port, cfg.Email.User, cfg.Web.AppPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendInvoiceTemplate(to string, data InvoiceData, templatePath string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	// Generate PDF
	pdfFile, err := GenerateInvoicePDF(data)
	if err != nil {
		fmt.Printf("Failed to generate PDF: %v\n", err)
		// Fallback to sending without attachment? Or return error?
		// User asked for file, so let's log and continue, or just return error.
		// For robustness, I'll return the error.
		return err
	}
	defer os.Remove(pdfFile) // Clean up

	return SendEmail(to, "Invoice Pembayaran - Kelas Santai", body.String(), pdfFile)
}
