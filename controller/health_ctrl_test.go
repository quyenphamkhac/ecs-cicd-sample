package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("200OK", func(t *testing.T) {
		r := gin.New()
		r.GET("/health", NewHealthController().Health)
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/health", nil)
		assert.NoError(t, err)
		r.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)
	})

	t.Run("SuccessResponse", func(t *testing.T) {
		r := gin.New()
		r.GET("/health", NewHealthController().Health)
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/health", nil)
		assert.NoError(t, err)
		r.ServeHTTP(rr, req)

		expected, err := json.Marshal(gin.H{
			"message": "This is blue/green deployment1",
		})

		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, expected, rr.Body.Bytes())
	})
}
