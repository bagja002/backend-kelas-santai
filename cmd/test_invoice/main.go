package main

import (
	"log"
	"project-kelas-santai/internal/config"
	"project-kelas-santai/internal/services"
	"time"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize email service
	emailService := services.NewEmailService(cfg)

	// Prepare data
	data := services.InvoiceData{
		InvoiceNumber: "INV-20241027-001",
		UserName:      "Barja Faskan",
		Date:          time.Now().Format("02 January 2006"),
		CourseName:    "Belajar Golang dari Nol",
		MentorName:    "Eko Kurniawan",
		Level:         "Intermediate",
		Total:         "1.500.000",
		GroupLink:     "https://t.me/kelas_santai_golang",
	}

	// Send invoice template
	templatePath := "/Users/ferdiansyah/Downloads/BackupProjec/project-kelas-santai/internal/templates/invoice.html"
	to := "barjafaskan9@gmail.com"

	log.Printf("Sending invoice to %s...", to)
	err = emailService.SendInvoiceTemplate(to, data, templatePath)
	if err != nil {
		log.Fatalf("Failed to send invoice: %v", err)
	}

	log.Println("Invoice sent successfully!")
}
