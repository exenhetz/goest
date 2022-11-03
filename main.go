package main

import (
	response "goesp/helpers"
	"goesp/modules"

	"github.com/labstack/echo/v4"
)

// Root route handler
func handleRootRoute(e echo.Context) error {
	return response.Success(e, map[string]string{
		"goest": "Hi :)",
	})
}

// The main function boots the whole app
func main() {

	// Initialize echo
	var app = echo.New()

	// Handle root route
	app.GET("/", handleRootRoute)

	// Prefix routes
	var v1 = app.Group("/api/v1")

	// Bind all module routes
	modules.BindModulesRoutes(v1)

	// Boot the server in port http://localhost:8099
	app.Logger.Fatal(app.Start(":8099"))
}
