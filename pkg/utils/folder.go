package utils

import (
	"fmt"
	"os"
)

func CreateFolder() {
	folders := []string{
		"public/testing",
	}

	for _, folder := range folders {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", folder, err)
			return
		}

	}

}
