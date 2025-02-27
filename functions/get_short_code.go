package functions

import "fmt"

// Function to convert QR / RFID code to shortcode
func GetShortCode(code *string) *string {
	if code == nil || *code == "" {
		return nil
	}
	firstChar := (*code)[0:1]
	lastSixChar := (*code)[len(*code)-6:]
	shortCode := fmt.Sprintf("%s*%s", firstChar, lastSixChar)
	return &shortCode
}
