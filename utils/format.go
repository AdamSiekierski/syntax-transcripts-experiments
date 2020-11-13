package utils

import (
	"encoding/json"
)

// Format - pretty-format given struct
func Format(data interface{}) string {
	formatted, _ := json.MarshalIndent(data, "", "  ")

	return string(formatted)
}
