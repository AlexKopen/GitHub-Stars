package main

const (
	ServerAddress              string = ":8080"
	StarGazerRequestEndpoint   string = "/stars"
	OwnerRepoSeparator         string = "/"
	OwnerRepoStringSplitLength int    = 2
)

const (
	GitHubErrorMessage           string = "There was an error fetching this repository information from GitHub. Make sure the repository exists and the owner and repo name is correct."
	RepositoryFormatErrorMessage string = "Invalid repository format.  Name must contain an owner and repo separated by a '/'."
	JSONFormatErrorMessage       string = "Invalid JSON format."
)
