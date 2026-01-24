package main

import (
	"fmt"
	"project-kelas-santai/internal/services"
)

func main() {

	invoiceData := services.InvoiceData{
		InvoiceNumber: "INV-123456",
		UserName:      "John Doe",
		Date:          "2023-01-01",
		CourseName:    "Golang for Beginners",
		MentorName:    "Mentor X",
		Level:         "Beginner",
		Total:         "100000",
		GroupLink:     "https://t.me/kelassantai",
	}

	// Use invoiceData to avoid "unused variable" error if necessary, or just definition.

	//Send Email asynchronously

	// Pastikan path template sesuai lokasi di server/local
	templatePath := "internal/templates/invoice.html"
	if err := services.SendInvoiceTemplate("barjafaskan9@gmail.com", invoiceData, templatePath); err != nil {
		fmt.Printf("Failed to send invoice email to %s: %v\n", "barjafaskan9@gmail.com", err)
	} else {
		fmt.Printf("Invoice sent to %s\n", "barjafaskan9@gmail.com")
	}

}
