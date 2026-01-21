package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type InvoiceData struct {
	InvoiceNumber string
	Date          string
	UserName      string
	UserEmail     string
	CourseName    string
	Level         string
	MentorName    string
	Price         string
	Total         string
	GroupLink     string
}

func main() {
	// Define paths
	cwd, _ := os.Getwd()
	templatePath := filepath.Join(cwd, "internal", "templates", "invoice.html")
	outputPath := filepath.Join(cwd, "test_invoice.html")

	// Parse template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Mock data
	data := InvoiceData{
		InvoiceNumber: "20241025-001",
		Date:          "25 Oktober 2024",
		UserName:      "Budi Santoso",
		UserEmail:     "budi@example.com",
		CourseName:    "Belajar Golang dari Nol",
		Level:         "Pemula",
		MentorName:    "Eko Kurniawan",
		Price:         "1.500.000",
		Total:         "1.500.000",
		GroupLink:     "https://t.me/kelassantai_golang",
	}

	// Create output file
	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer f.Close()

	// Execute template
	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	log.Printf("Successfully created test invoice at: %s", outputPath)
}
