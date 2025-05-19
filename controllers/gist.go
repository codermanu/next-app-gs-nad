package controllers

import (
	"net/http"
	"portal/register/model" // Import the models package
	// No need to import the service package directly, use apiCtx.Service

	"github.com/cloudimpl/next-coder-sdk/apicontext"
	"github.com/cloudimpl/next-coder-sdk/polycode"
	"github.com/gin-gonic/gin"
)

// SaveGist handles the HTTP request to save a markdown file as a GitHub Gist.
func SaveGist(c *gin.Context) {
	request := model.SaveGistRequest{}

	// Bind the JSON request body to the request struct
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

	// Call the "gist-service" and its "SaveMarkdownGist" function
	// The service name "gist-service" must match the folder name under services/
	// The function name "SaveMarkdownGist" must match the function name in the service file
	// RequestReply sends the request to the service and waits for a reply
	// GetAny() retrieves the response object (which should be model.SaveGistResponse)
	resp, err := apiCtx.Service("gist-service").Get().
		RequestReply(polycode.TaskOptions{}, "SaveMarkdownGist", request).GetAny()

	if err != nil {
		// Handle errors returned by the service function
		// In a real application, you might inspect the error type (e.g., using polycode.IsError)
		// to return different HTTP status codes (e.g., 400 for validation errors, 500 for internal errors)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the successful response with HTTP status 201 Created
	c.JSON(http.StatusCreated, resp)
}