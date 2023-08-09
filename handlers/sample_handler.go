package handlers

import (
	_ "com/adl/et/telco/dte/template/baseapp/docs"
	"com/adl/et/telco/dte/template/baseapp/dtos"
	"com/adl/et/telco/dte/template/baseapp/services"
	"com/adl/et/telco/dte/template/baseapp/utilities"
	"com/adl/et/telco/dte/template/baseapp/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IntergrationHandler is an API endpoint that processes a sample request and returns a response.
// @Summary Process a sample request
// @Description This endpoint processes a sample request and returns a response.
// @Tags Basic Template
// @Accept json
// @Produce json
// @Param request body dtos.SampleRequestEntity true "Sample Request Body"
// @Success 200 {object} dtos.SampleResponseEntity
// @Failure 400 {object} utilities.ErrorResponse
// @Failure 404 {object} utilities.ErrorResponse
// @Failure 500 {object} utilities.ErrorResponse
// @Router /action [post]
func IntergrationHandler(c *gin.Context) {
	var sampleRequestEntity dtos.SampleRequestEntity
	if err := c.ShouldBindJSON(&sampleRequestEntity); err != nil {
		c.JSON(http.StatusBadRequest, utilities.ErrorResponse{Message: "invalid request payload"})
		return
	}

	//validate the request
	if err := validators.Validate(sampleRequestEntity); err != nil {
		c.JSON(http.StatusBadRequest, utilities.ErrorResponse{Message: err.Error()})
		return

	}
	// Call domain business logic
	sampleResponseEntity := services.Process(sampleRequestEntity)

	// Return the response
	switch sampleResponseEntity.ResCode {
	case "200":
		c.JSON(http.StatusOK, sampleResponseEntity)
	case "202":
		c.JSON(http.StatusAccepted, sampleResponseEntity)
	case "400":
		c.JSON(http.StatusBadRequest, sampleResponseEntity)
	case "404":
		c.JSON(http.StatusNotFound, sampleResponseEntity)
	default:
		c.JSON(http.StatusInternalServerError, sampleResponseEntity)

	}
}
