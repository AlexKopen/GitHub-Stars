package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
	// CLI arguments to suppress error messages.  False by default.
	suppressErrorsPtr := flag.Bool("suppress", false, "Suppressing client errors")
	flag.Parse()
	suppressErrors := *suppressErrorsPtr

	// Create a new buffered reader for user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(WelcomeMessage)

	// Iterate infinitely to process each individual input
	for {
		fmt.Printf("\n%s\n", InstructionMessage)
		// Process the user's input after enter is pressed, stripping any new lines
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		// Send the input off for validation
		validInputs := generateValidInputs(userInput, suppressErrors)

		// Process any available valid inputs
		if len(validInputs) > 0 {
			output := processInput(validInputs)

			// Print the results
			if len(output) > 0 {
				fmt.Println("\nRESULTS:")
				fmt.Printf("\n%s\n", output)
			}
		}
	}
}

func generateValidInputs(input string, suppressErrors bool) []string {
	// Trim leading and trailing whitespace
	input = strings.TrimSpace(input)

	// Split the input by commas
	splitInputs := strings.Split(input, ",")

	// Only process valid inputs
	var validInputs []string

	// For each repository, determine whether the format is valid
	for _, repository := range splitInputs {
		// Check to make sure the input contains a slash
		repository = strings.TrimSpace(repository)
		splitRepositoryName := strings.Split(repository, "/")
		if len(splitRepositoryName) != 2 {
			if !suppressErrors {
				fmt.Printf("\n%s %s.  %s\n", InvalidRepo, repository, InvalidRepoSlashError)
			}
			continue
		}

		// Check to make sure input is entered for both the organization and repository
		splitRepositoryName[0] = strings.TrimSpace(splitRepositoryName[0])
		splitRepositoryName[1] = strings.TrimSpace(splitRepositoryName[1])

		if len(splitRepositoryName[0]) < 1 || len(splitRepositoryName[1]) < 1 {
			if !suppressErrors {
				fmt.Printf("\n%s %s.  %s\n", InvalidRepo, repository, InvalidRepoNameError)
			}
			continue
		}

		// If all conditions have passed, add the repository to the list of valid inputs
		validInputs = append(validInputs, repository)
	}

	return validInputs
}

func processInput(validInputs []string) string {
	// Create a new request payload
	starGazerRequest := StarGazerRequest{Repositories: validInputs}

	requestBodyJSON, _ := json.Marshal(starGazerRequest)

	//	Call the API running on the server
	resp, serverRequestErr := http.Post(ServerStarsEndpointURL, "application/json", bytes.NewBuffer(requestBodyJSON))
	if serverRequestErr != nil {
		fmt.Printf("\n%s\n", ServerRequestError)
		return ""
	}

	// Read the response data
	responseData, bodyReadErr := ioutil.ReadAll(resp.Body)
	if bodyReadErr != nil {
		fmt.Printf("\n%s\n", ServerParseError)
		return ""
	}

	// Convert the response body to a star gazer response struct
	starGazerResponse := StarGazerResponse{}
	responseParsingError := json.Unmarshal(responseData, &starGazerResponse)
	if responseParsingError != nil {
		fmt.Printf("\n%s\n", ResponseParseError)
		return ""
	}

	// There are no valid results, don't return the response
	if len(starGazerResponse.StarGazerResults) == 0 {
		return ""
	}

	// Format the output
	formattedOutput, _ := json.MarshalIndent(starGazerResponse, "", "\t")
	return string(formattedOutput)
}
