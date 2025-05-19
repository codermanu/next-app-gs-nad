package gist_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"portal/register/error-codes" // Corrected import path
	"portal/register/model"

	"github.com/cloudimpl/next-coder-sdk/polycode"
)

// Placeholder for GitHub Token - MUST be handled securely in production
// Consider using environment variables, a secrets manager, or Polycode configuration features
const githubToken = "YOUR_GITHUB_TOKEN" // Replace with secure retrieval mechanism

// SaveMarkdownGist saves the provided markdown content as a GitHub Gist.
func SaveMarkdownGist(ctx polycode.ServiceContext, req model.SaveGistRequest) (model.SaveGistResponse, error) {
	if githubToken == "YOUR_GITHUB_TOKEN" {
		// Use the defined error code
		return model.SaveGistResponse{}, error_codes.GitHubTokenNotConfigured.With().Wrap(fmt.Errorf("GitHub token is not configured"))
	}

	// Construct the request body for the GitHub Gist API
	gistContent := map[string]interface{}{
		"description": req.Description,
		"public":      req.Public,
		"files": map[string]map[string]string{
			req.Filename: {
				"content": req.Content,
			},
		},
	}

	body, err := json.Marshal(gistContent)
	if err != nil {
		// Wrap standard errors with a Polycode error if appropriate, or return as is
		return model.SaveGistResponse{}, fmt.Errorf("failed to marshal gist content: %w", err)
	}

	// Create HTTP client and request
	client := &http.Client{}
	httpReq, err := http.NewRequest("POST", "https://api.github.com/gists", bytes.NewBuffer(body))
	if err != nil {
		return model.SaveGistResponse{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Add required headers, including authentication
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "token "+githubToken)
	httpReq.Header.Set("Accept", "application/vnd.github.v3+json") // Recommended by GitHub API docs

	// Perform the request
	resp, err := client.Do(httpReq)
	if err != nil {
		return model.SaveGistResponse{}, fmt.Errorf("failed to perform HTTP request to GitHub: %w", err)
	}
	defer resp.Body.Close()

	// Read response body for potential errors or success details
	var responseBody bytes.Buffer
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		return model.SaveGistResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check the HTTP status code
	if resp.StatusCode != http.StatusCreated {
		// Use the defined Polycode error for API failures
		// Include status code and response body in the error details
		return model.SaveGistResponse{}, error_codes.GitHubAPIError.With(resp.StatusCode, responseBody.String()).Wrap(fmt.Errorf("GitHub API returned status %d", resp.StatusCode))
	}

	// Parse the successful response body to get the Gist URL and ID
	var gistResp struct {
		HTMLURL string `json:"html_url"`
		ID      string `json:"id"`
	}
	err = json.Unmarshal(responseBody.Bytes(), &gistResp)
	if err != nil {
		return model.SaveGistResponse{}, fmt.Errorf("failed to parse GitHub response: %w", err)
	}

	// Return the successful response DTO
	return model.SaveGistResponse{
		GistURL: gistResp.HTMLURL,
		GistID:  gistResp.ID,
	}, nil
}