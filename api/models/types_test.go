package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswerPairJSONMarshalling(t *testing.T) {
	pair := AnswerPair{
		Question: "What is Go?",
		Answer:   "A programming language by Google.",
	}

	// Marshal to JSON
	data, err := json.Marshal(pair)
	assert.NoError(t, err)

	expectedJSON := `{"question":"What is Go?","answer":"A programming language by Google."}`
	assert.JSONEq(t, expectedJSON, string(data))

	var decoded AnswerPair
	err = json.Unmarshal(data, &decoded)
	assert.NoError(t, err)

	assert.Equal(t, pair, decoded)
}
