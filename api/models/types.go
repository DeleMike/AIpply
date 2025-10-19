// Package models used to store app-wide data objects
package models

// AnswerPair is a struct used to help AI view how a related answer is related to a question
type AnswerPair struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
