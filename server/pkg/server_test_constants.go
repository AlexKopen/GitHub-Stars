package main

const (
	SuccessfulRequest = `
		{
  			"Repositories": [
    			"AlexKopen/AUR-Package-Scraper",
    			"angular/angular"
		  	]
		}
	`
	InvalidPayloadRequest = `
		{
  			"Typo": [
    			"AlexKopen/AUR-Package-Scraper"
		  	]
		}
	`
	RepositoryFormatErrorMessageRequest = `
		{
  			"Repositories": [
		    	"Invalid"
		  	]
		}
	`

	NonExistentGithubRepoErrorMessagesRequest = `
		{
  			"Repositories": [
		    	"doesnot/exist123"
		  	]
		}
	`
)
