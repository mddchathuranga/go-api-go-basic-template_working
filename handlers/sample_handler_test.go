package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIntergrationHandler(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Define a route with the IntergrationHandler
	router.POST("/action", IntergrationHandler)

	// Create a sample request JSON payload
	requestPayload := `{"id": 1}`

	// Create a new HTTP request with the sample payload
	req, err := http.NewRequest("POST", "/action", bytes.NewBufferString(requestPayload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the Gin router to handle the request
	router.ServeHTTP(rr, req)

	// Assert the response status code is 200 (OK)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert the response body contains the expected message
	expectedResponse := `{"resCode":"200","resDesc":"Operation Success","errorMessage":""}`
	assert.Equal(t, expectedResponse, rr.Body.String())
}
