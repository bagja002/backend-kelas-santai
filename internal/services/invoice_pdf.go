package services

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

// GenerateInvoicePDF creates a PDF invoice based on the provided data
// and returns the path to the temporary file.
func GenerateInvoicePDF(data InvoiceData) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// New Palette
	// Primary (Brown): #5D4037 -> 93, 64, 55
	primaryColor := []int{93, 64, 55}
	// Accent (Light Brown): #8D6E63 -> 141, 110, 99
	// Text (Dark Brown): #3E2723 -> 62, 39, 35
	textColor := []int{62, 39, 35}

	// Header Background
	pdf.SetFillColor(primaryColor[0], primaryColor[1], primaryColor[2])
	pdf.Rect(0, 0, 210, 50, "F")

	// Header Text
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 24)
	pdf.SetXY(10, 15)
	pdf.Cell(0, 10, "Kelas Santai.")

	pdf.SetFont("Arial", "", 12)
	pdf.SetXY(10, 25)
	pdf.SetTextColor(239, 235, 233) // #EFEBE9 Light beige for subtext
	pdf.Cell(0, 10, fmt.Sprintf("#%s", data.InvoiceNumber))

	// Success Message
	pdf.SetTextColor(textColor[0], textColor[1], textColor[2])
	pdf.SetFont("Arial", "B", 16)
	pdf.SetXY(10, 60)
	pdf.Cell(0, 10, "Pembayaran Berhasil!")

	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(141, 110, 99) // Accent brown
	pdf.SetXY(10, 70)
	pdf.Cell(0, 10, fmt.Sprintf("Terima kasih, %s! Akses kelasmu sudah aktif.", data.UserName))

	// Invoice Details
	yStart := 90.0
	pdf.SetDrawColor(239, 235, 233) // Light beige border
	pdf.SetLineWidth(0.5)

	// Helper to draw rows
	drawRow := func(label, value string, isBold bool, valueColor []int) {
		pdf.SetTextColor(141, 110, 99) // Label Color (Accent)
		pdf.SetFont("Arial", "", 11)
		pdf.SetXY(10, yStart)
		pdf.Cell(50, 10, label)

		if isBold {
			pdf.SetFont("Arial", "B", 11)
		} else {
			pdf.SetFont("Arial", "", 11)
		}

		if len(valueColor) == 3 {
			pdf.SetTextColor(valueColor[0], valueColor[1], valueColor[2])
		} else {
			pdf.SetTextColor(textColor[0], textColor[1], textColor[2]) // Default text color
		}

		pdf.SetXY(60, yStart)
		pdf.Cell(0, 10, value)

		// Line separator
		pdf.Line(10, yStart+10, 200, yStart+10)
		yStart += 12
	}

	drawRow("Tanggal", data.Date, false, nil)
	drawRow("Kelas", data.CourseName, true, primaryColor)
	drawRow("Mentor", data.MentorName, false, nil)
	drawRow("Level", data.Level, false, nil)

	// Total Row
	yStart += 5
	pdf.SetFillColor(239, 235, 233) // #EFEBE9 Light beige bg
	pdf.Rect(10, yStart, 190, 15, "F")

	pdf.SetTextColor(textColor[0], textColor[1], textColor[2])
	pdf.SetFont("Arial", "B", 12)
	pdf.SetXY(15, yStart+2.5)
	pdf.Cell(50, 10, "Total Pembayaran")

	pdf.SetTextColor(primaryColor[0], primaryColor[1], primaryColor[2])
	pdf.SetFont("Arial", "B", 14)
	pdf.SetXY(60, yStart+2.5)
	pdf.Cell(0, 10, fmt.Sprintf("Rp %s", data.Total))

	// Footer
	pdf.SetFillColor(textColor[0], textColor[1], textColor[2]) // Dark footer
	pdf.Rect(0, 260, 210, 37, "F")

	pdf.SetY(270)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(239, 235, 233) // Beige text
	pdf.CellFormat(0, 10, "Kelas Santai - Jakarta, Indonesia", "", 0, "C", false, 0, "")

	// Save
	fileName := fmt.Sprintf("invoice-%s.pdf", data.InvoiceNumber)
	if err := pdf.OutputFileAndClose(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}
