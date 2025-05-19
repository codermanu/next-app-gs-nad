package main

import (
	"github.com/cloudimpl/next-coder-sdk/polycode"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "portal/register/.polycode" // Import for side effects (service registration)
	"portal/register/controllers"
	"portal/register/lib"
	// Import the new controller package if it's in a subfolder like controllers/gist
	// "portal/register/controllers/gist" // If gist controller is in a subfolder
)

func main() {
	// Initialize and set validator for Polycode runtime
	v := lib.NewValidator()
	polycode.SetValidator(v)

	// Setup Gin router and CORS middleware
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// Add any custom headers used by polycode or your application
	config.AllowHeaders = append(config.AllowHeaders, "x-polycode-partition-key")
	r.Use(cors.New(config))

	// Register controller endpoints
	// Existing endpoint (using api.FromWorkflow pattern)
	r.POST("/greeting", controllers.Greeting)

	// New Gist endpoint (using the *gin.Context pattern)
	// Assuming SaveGist is in controllers/gist.go or controllers/greeting.go (if combined)
	// If controllers are split into sub-packages, import and use them like:
	// r.POST("/gists", gist.SaveGist) // If SaveGist is in controllers/gist.go
	// For this example, we assume controllers are in the main controllers package
	r.POST("/gists", controllers.SaveGist)

	// Start the Polycode application
	polycode.StartApp(r)
}