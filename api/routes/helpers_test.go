package routes

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DeleMike/AIpply/api/apierrors"
	"github.com/DeleMike/AIpply/api/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genai"
)

func TestHandleLLMRequestPayload(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("invalid JSON returns 400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{invalid-json"))
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		handleLLMRequest(c, &struct{}{}, func(ctx context.Context, _ *struct{}) (any, error) {
			return nil, nil
		})
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), apierrors.ErrInvalidRequest)
	})

	t.Run("LLM client not initialized returns 500", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		// Ensure no client
		service.LLMClient = nil

		handleLLMRequest(c, &struct{}{}, func(ctx context.Context, _ *struct{}) (any, error) {
			return nil, nil
		})

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), apierrors.ErrLLMNotInitialized)
	})
	t.Run("handler error returns 500", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		// Mock client to pass readiness check
		service.LLMClient = &genai.Client{}

		handleLLMRequest(c, &struct{}{}, func(ctx context.Context, _ *struct{}) (any, error) {
			return nil, errors.New("failed")
		})

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), apierrors.ErrWeCouldNotProcessRequest)
	})
	t.Run("success returns 200", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		service.LLMClient = &genai.Client{}

		handleLLMRequest(c, &struct{}{}, func(ctx context.Context, _ *struct{}) (any, error) {
			return gin.H{"message": "ok"}, nil
		})

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "ok")
	})
}
