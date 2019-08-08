package server

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestNotFoundError(t *testing.T) {
	writer:=httptest.NewRecorder()
	request:=httptest.NewRequest("GET","/hello",nil)
	NotFoundError(writer,request)
	assert.Equal(t,"{\"code\":\"404\",\"message\":\"service not found\"}",writer.Body.String())
}
