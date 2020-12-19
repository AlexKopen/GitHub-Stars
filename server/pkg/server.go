package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"strings"
	"sync"
)

type StarGazerRequest struct {
	Repositories []string
}

type StarGazerResult struct {
	Repository   string
	Count        int
	ErrorMessage string
}

type StarGazerResponse struct {
	TotalCount       int
	StarGazerResults []StarGazerResult
}

func main() {
	// Run the API on port 8080
	_ = initializeServer().Run(ServerAddress)
}

func initializeServer() *gin.Engine {
	// Set up a default gin router
	router := gin.Default()

	// Initiate the /stars endpoint
	router.POST(StarGazerRequestEndpoint, starGazerPOST)

	return router
}

func starGazerPOST(c *gin.Context) {
	// Bind the request payload to a JSON struct
	var starGazerRequest StarGazerRequest
	err := c.BindJSON(&starGazerRequest)

	// If there's an error or the JSON doesn't contain a repositories array, throw an error
	if err != nil || len(starGazerRequest.Repositories) == 0 {
		c.JSON(400, gin.H{"Error": JSONFormatErrorMessage})
	} else {
		// No errors, begin star gazer processing
		processRepositories(c, starGazerRequest)
	}
}

func processRepositories(c *gin.Context, request StarGazerRequest) {
	// Create a wait group to manage goroutines
	var wg sync.WaitGroup
	// Initialize a response struct
	var response StarGazerResponse
	// Create a go-github client to interface with the GitHub API
	client := github.NewClient(nil)

	// Process each <organization>/<repository> string concurrently
	for _, repoName := range request.Repositories {
		wg.Add(1)
		go findStarCount(repoName, client, &response, &wg)
	}

	// Don't finish execution until all repositories have been processed
	wg.Wait()

	// Send a success with the star gazer response struct as JSON upon completion
	c.JSON(200, response)
}

func findStarCount(repository string, client *github.Client, response *StarGazerResponse, wg *sync.WaitGroup) {
	// Initialize a star gazer result struct to be appended to the final response struct
	var result StarGazerResult
	result.Repository = repository

	// Extract the owner and repository name
	ownerRepoSplit := strings.Split(repository, OwnerRepoSeparator)

	// Only process if there's an owner and repo name to work with
	if len(ownerRepoSplit) == OwnerRepoStringSplitLength {
		repo, _, err := client.Repositories.Get(context.Background(), ownerRepoSplit[0], ownerRepoSplit[1])

		// If the repo is not found, assign an error to the result error message attribute
		if err != nil {
			result.ErrorMessage = GitHubErrorMessage
		} else {
			// Repo is found, assign the star gazer count
			result.Count = *repo.StargazersCount
		}

	} else {
		// Repository name invalid, assign an error to the result error message attribute
		result.ErrorMessage = RepositoryFormatErrorMessage
	}

	// Increase the total count and append the newly constructed result to the final response struct
	response.TotalCount += result.Count
	response.StarGazerResults = append(response.StarGazerResults, result)

	// Mark this goroutine as complete
	wg.Done()
}
