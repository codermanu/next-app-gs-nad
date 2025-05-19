package main

import (
	"github.com/cloudimpl/next-coder-sdk/polycode"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "portal/register/.polycode" // Import for side effects (service registration)
	"portal/register/controllers"
	"portal/register/lib"
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
	// Greeting endpoint using the standard Gin controller pattern
	r.POST("/greeting", controllers.HandleGreeting)

	// Gist endpoint using the standard Gin controller pattern
	r.POST("/gists", controllers.SaveGist)

	// Start the Polycode application
	polycode.StartApp(r)
}