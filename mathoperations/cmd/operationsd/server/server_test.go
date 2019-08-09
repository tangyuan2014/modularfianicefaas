package server

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOperation(t *testing.T) {
	req,err:=http.NewRequest("GET","/?number=3",nil)
	if err!=nil{
		log.Fatal(err.Error())}
	resp:=httptest.NewRecorder()
	Operation(resp,req)
	expectresp:="result is 6"
	assert.Equal(t,http.StatusOK,resp.Code)
	assert.Equal(t,expectresp,resp.Body.String())
}

func TestOperation2(t *testing.T) {
	req,err:=http.NewRequest("GET","/?number=a",nil)
	if err!=nil{
		log.Fatal(err.Error())}
	resp:=httptest.NewRecorder()
	Operation(resp,req)
	expectresp:=`{"Code":400,"Message":"Please provide a integer "}`
	assert.Equal(t,http.StatusBadRequest,resp.Code)
	assert.Equal(t,expectresp,resp.Body.String())
}
