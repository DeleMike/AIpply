// Package utils contains utility functions and error handling mechanisms for the AIpply application.
package utils

// MaskString hides sensitve strings info
func MaskString(s string) string {
	if len(s) <= 4 {
		return "****"
	}
	return s[:2] + "****" + s[len(s)-2:]
}
