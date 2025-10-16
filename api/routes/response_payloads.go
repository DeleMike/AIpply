package routes

type QuestionsResponsePayload struct {
	Message   string   `json:"message"`
	Questions []string `json:"questions"`
}
