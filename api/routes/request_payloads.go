// Package routes provides the HTTP route definitions for the AIpply.
package routes

// GenerateQuestionsRequest represents the expected payload for generating interview questions.
type GenerateQuestionsRequest struct {
	Name            string `json:"name" binding:"required"`
	JobDescription  string `json:"job_description" binding:"required"`
	ExperienceLevel string `json:"experience_level" binding:"required"`
}
