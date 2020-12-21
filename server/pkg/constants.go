package main

const (
	ServerAddress              string = ":8080"
	StarGazerRequestEndpoint   string = "/api/v1/stars"
	OwnerRepoSeparator         string = "/"
	OwnerRepoStringSplitLength int    = 2
)

const (
	RepositoryFormatErrorMessage string = "Invalid repository format.  Name must contain an owner and repo separated by a '/'."
	GitHubErrorMessage           string = "There was an error fetching this repository information from GitHub. Make sure the repository exists and the owner and repo name is correct."
	JSONFormatErrorMessage       string = "Invalid JSON format in request."
)
