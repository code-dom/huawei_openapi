package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

// Convert image to Base64 string
func EncodeImageToBase64(imagePath string) (string, error) {
	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(imageBytes), nil
}

func DecodeBase64(encodedStr string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encodedStr)
}

func DecodeBase64ToFile(encodedStr string, target string) error {
	decoded, err := DecodeBase64(encodedStr)
	if err != nil {
		return fmt.Errorf("decode base64 failed: %w", err)
	}
	err = os.WriteFile(target, decoded, 0644)
	if err != nil {
		return fmt.Errorf("write file failed: %w", err)
	}
	return nil
}
