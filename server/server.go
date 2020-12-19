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

	router.POST("/stars", starGazerPOST)

	_ = router.Run(":8080")
}

func starGazerPOST(c *gin.Context) {
	var starGazerRequest StarGazerRequest
	_ = c.BindJSON(&starGazerRequest)

	processRepositories(c, starGazerRequest)
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

	ownerRepoSplit := strings.Split(repository, "/")

	if len(ownerRepoSplit) == 2 {
		repo, _, err := client.Repositories.Get(context.Background(), ownerRepoSplit[0], ownerRepoSplit[1])

		if err != nil {
			result.ErrorMessage = "There was an error fetching this repository information from GitHub"
		} else {
			result.Count = *repo.StargazersCount
		}

	} else {
		result.ErrorMessage = "Invalid repository format.  Name must contain an owner and repo separated by a '/'"
	}

	response.TotalCount += result.Count
	response.StarGazerResults = append(response.StarGazerResults, result)

	wg.Done()
}
