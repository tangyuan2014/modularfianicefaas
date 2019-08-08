package server

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestNotFoundError(t *testing.T) {
	writer := httptest.NewRecorder()
	NotFoundError(writer)
	assert.Equal(t, "{\"code\":\"404\",\"message\":\"service not found\"}", writer.Body.String())
}
