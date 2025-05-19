package controllers

import (
	"net/http" // Import net/http for status codes
	"portal/register/model"

	"github.com/cloudimpl/next-coder-sdk/apicontext"
	"github.com/cloudimpl/next-coder-sdk/polycode"
	"github.com/gin-gonic/gin"
)

// HandleGreeting is a Gin controller function to handle the /greeting endpoint.
// It calls the "greeting-service" to generate the greeting message.
func HandleGreeting(c *gin.Context) {
	// Bind the JSON request body to the request struct
	request := model.HelloRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the apiContext from the request context
	apiCtx, err := apicontext.FromContext(c.Request.Context())
	if err != nil {
		// This error indicates a problem with the Polycode runtime context setup
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get api context"})
		return
	}

	// Call the "greeting-service" and its "Greeting" function
	// The service name "greeting-service" must match the folder name under services/
	// The function name "Greeting" must match the function name in the service file (services/greeting-service/Greeting.go)
	// RequestReply sends the request to the service and waits for a reply
	// Get() retrieves the response object and attempts to unmarshal it into the provided variable
	var output model.HelloResponse
	err = apiCtx.Service("greeting-service").Get().
		RequestReply(polycode.TaskOptions{}, "Greeting", request).Get(&output)

	if err != nil {
		// Handle errors returned by the service function
		// In a real application, you might inspect the error type (e.g., using polycode.IsError)
		// to return different HTTP status codes (e.g., 400 for validation errors, 500 for internal errors)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the successful response with HTTP status 200 OK
	c.JSON(http.StatusOK, output)
}