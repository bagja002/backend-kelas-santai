package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// HandleSingleFileUpload handles the upload of a single file
func HandleSingleFileUpload(c *fiber.Ctx, paramName, destFolder string) (string, error) {
	// Get the file from the form
	file, err := c.FormFile(paramName)
	if err != nil {
		return "", err
	}

	// Validate file extension (optional, can be expanded)
	// allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	// ... validation logic

	// Create destination folder if it doesn't exist
	if err := os.MkdirAll(destFolder, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// Check if original filename had spaces, maybe sanitize it, but timestamp is safe enough
	// Clean relative path to prevent directory traversal
	cleanDest := filepath.Clean(destFolder)
	fullPath := filepath.Join(cleanDest, filename)

	// Save the file
	if err := c.SaveFile(file, fullPath); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	// Return the relative path for storage
	// using forward slashes for URL compatibility
	return strings.ReplaceAll(filepath.Join(destFolder, filename), "\\", "/"), nil
}
