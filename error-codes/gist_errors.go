package error_codes

import "github.com/cloudimpl/next-coder-sdk/polycode"

// Module name should be unique for the microservice
const moduleName = "portal.register" // Assuming "portal/register" is the module name

// Gist related error codes starting from 1000
var GitHubAPIError = polycode.DefineError(moduleName, 1000, "GitHub API Error: Status %d, Response: %s")
var GitHubTokenNotConfigured = polycode.DefineError(moduleName, 1001, "GitHub token is not configured")