package main

const (
	WelcomeMessage         string = "GitHub Stargazer application started"
	InstructionMessage     string = "Enter a list of <organization>/<repository> inputs, separated by commas, then press Enter. Example - angular/angular, twilio/twilio-python:"
	InvalidRepo            string = "Invalid repository name:"
	InvalidRepoSlashError  string = "Name must contain exactly one '/'.  This input will not be processed."
	InvalidRepoNameError   string = "The organization and repository must contain valid alphanumeric characters.  This input will not be processed."
	ResponseParseError     string = "There was an error parsing the response from the server."
	ServerStarsEndpointURL string = "http://localhost:8080/stars"
	ServerRequestError     string = "There was an error requesting data from the server."
	ServerParseError       string = "There was an error parsing the response from the server."
)
