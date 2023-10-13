package util

import (
	"strings"

	"github.com/google/uuid"
)

func NormalizeEmail(email string) string {
	return strings.ToLower(email)
}

func NewID() string {
	return uuid.New().String()
}

func Slug(text string) string {
	return strings.ToLower(strings.ReplaceAll(text, " ", "-"))
}
