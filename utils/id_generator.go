package utils

import "github.com/google/uuid"

// GenerateUUID creates a new unique identifier
func GenerateUUID() string {
	return uuid.New().String()
}
