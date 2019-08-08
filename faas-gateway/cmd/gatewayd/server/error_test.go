package server

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFoundError(t *testing.T) {
	writer := httptest.NewRecorder()
	NotFoundError(writer)
	assert.Equal(t, "{\"code\":\"404\",\"message\":\"service not found\"}", writer.Body.String())
}
