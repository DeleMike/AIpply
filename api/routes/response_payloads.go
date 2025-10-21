package routes

// QuestionsResponsePayload represents the structure of the JSON response
// returned when interview questions are successfully generated.
type QuestionsResponsePayload struct {
	Message   string   `json:"message"`
	Questions []string `json:"questions"`
}

// CVResponsePayload represent the structure of response
// returned when the CV is ready
type CVResponsePayload struct {
	Message string `json:"message"`
	CV      string `json:"cv"`
}

// CoverLetterResponsePayload represent the structure of response
// returned when the cover letter is ready
type CoverLetterResponsePayload struct {
	Message string `json:"message"`
	Cover   string `json:"cover"`
}
