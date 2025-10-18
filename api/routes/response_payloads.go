package routes

// QuestionsResponsePayload represents the structure of the JSON response
// returned when interview questions are successfully generated.
type QuestionsResponsePayload struct {
	Message   string   `json:"message"`
	Questions []string `json:"questions"`
}

type CVResponsePayload struct {
	Message string `json:"message"`
	CV      string `json:"cv"`
}
