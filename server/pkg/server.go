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
	router := gin.Default()

	router.POST(StarGazerRequestEndpoint, starGazerPOST)

	_ = router.Run(ServerAddress)
}

func starGazerPOST(c *gin.Context) {
	var starGazerRequest StarGazerRequest
	err := c.BindJSON(&starGazerRequest)

	if err != nil || len(starGazerRequest.Repositories) == 0 {
		c.JSON(400, gin.H{"Error": JSONFormatErrorMessage})
	} else {
		processRepositories(c, starGazerRequest)
	}
}

func processRepositories(c *gin.Context, request StarGazerRequest) {
	var wg sync.WaitGroup
	var response StarGazerResponse

	client := github.NewClient(nil)

	for _, repoName := range request.Repositories {
		wg.Add(1)
		go findStarCount(repoName, client, &response, &wg)
	}

	wg.Wait()

	c.JSON(200, response)
}

func findStarCount(repository string, client *github.Client, response *StarGazerResponse, wg *sync.WaitGroup) {
	var result StarGazerResult
	result.Repository = repository

	ownerRepoSplit := strings.Split(repository, OwnerRepoSeparator)

	if len(ownerRepoSplit) == OwnerRepoStringSplitLength {
		repo, _, err := client.Repositories.Get(context.Background(), ownerRepoSplit[0], ownerRepoSplit[1])

		if err != nil {
			result.ErrorMessage = GitHubErrorMessage
		} else {
			result.Count = *repo.StargazersCount
		}

	} else {
		result.ErrorMessage = RepositoryFormatErrorMessage
	}

	response.TotalCount += result.Count
	response.StarGazerResults = append(response.StarGazerResults, result)

	wg.Done()
}
