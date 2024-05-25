package secretenv

import (
	"io"
	"os"
	"strings"
)

// Get first checks if the key suffixed with "_FILE" exists.
// If it does, it reads the content of the file and returns it.
// Otherwise, it returns the environment value of the key.
func Get(key string) string {
	if key == "" {
		return ""
	}

	fileKey := key + "_FILE"

	if value, ok := os.LookupEnv(fileKey); ok {
		file, err := os.Open(value)
		if err == nil {
			defer file.Close()

			data, err := io.ReadAll(file)
			if err == nil {
				return strings.TrimSpace(string(data))
			}
		}
	}

	return os.Getenv(key)
}
