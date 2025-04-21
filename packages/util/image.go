package util

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func SaveFile(file *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(file.Filename)
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}

	if !contains(allowedExtensions, ext) {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	fileName := fmt.Sprintf("%s_%s", generateRandomString(10), file.Filename)
	filePath := filepath.Join("storages", "images", fileName)

	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return "", err
	}

	return filePath, nil
}

// contains checks if a slice contains a specific string
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[seededRand.Intn(len(charset))])
	}
	return sb.String()
}
