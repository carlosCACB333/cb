package utils

import (
	"strings"
)

func NormalizeEmail(email string) string {
	return strings.ToLower(email)
}

func GenerateSlug(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}
