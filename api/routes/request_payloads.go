// Package routes provides the HTTP route definitions for the AIpply.
package routes

import "github.com/DeleMike/AIpply/api/models"

// GenerateQuestionsRequest represents the expected payload for generating interview questions.
type GenerateQuestionsRequest struct {
	JobDescription  string `json:"job_description" binding:"required"`
	ExperienceLevel string `json:"experience_level" binding:"required"`
}

// GenerateCVRequest represents the expected payload for generating a CV
type GenerateCVRequest struct {
	JobDescription string              `json:"job_description" binding:"required"`
	Answers        []models.AnswerPair `json:"answers" binding:"required"`
}
