// Package stringutil holds simple string extension methods
package stringutil

// MaskString hides sensitve strings info
func MaskString(s string) string {
	if len(s) <= 4 {
		return "****"
	}
	return s[:2] + "****" + s[len(s)-2:]
}
