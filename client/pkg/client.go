package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("GitHub Stargazer application started")

	for {
		fmt.Println("\nEnter a list of <organization>/<repository> inputs, separated by commas. Example - angular/angular, twilio/twilio-python:\n")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		inputValidation(userInput)

	}

}

func inputValidation(input string) {
	// Trim leading and trailing whitespace
	input = strings.TrimSpace(input)
	splitInputs := strings.Split(input, ",")

	var validInputs []string
	for _, repository := range splitInputs {
		repository = strings.TrimSpace(repository)
		splitRepositoryName := strings.Split(repository, "/")
		if len(splitRepositoryName) != 2 {
			fmt.Printf("\nInvalid repository name: %s.  Name must contain exactly one '/'.  This input will not be processed.\n", repository)
			continue
		}

		splitRepositoryName[0] = strings.TrimSpace(splitRepositoryName[0])
		splitRepositoryName[1] = strings.TrimSpace(splitRepositoryName[1])

		if len(splitRepositoryName[0]) < 1 || len(splitRepositoryName[1]) < 1 {
			fmt.Printf("\nInvalid repository name: %s.  The organization and repository must contain valid alphanumeric characters.  This input will not be processed.\n", repository)
			continue
		}

		validInputs = append(validInputs, repository)
	}

	if len(validInputs) > 0 {
		processInput(validInputs)
	}

}

func processInput(validInputs []string) {
	//	Call the API
	starGazerRequest := StarGazerRequest{Repositories: validInputs}

	url := "http://localhost:8080/stars"
	jsonValue, _ := json.Marshal(starGazerRequest)

	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

	responseData, bodyReadErr := ioutil.ReadAll(resp.Body)
	if bodyReadErr != nil {
		log.Fatal(bodyReadErr)
	}

	starGazerResponse := StarGazerResponse{}
	_ = json.Unmarshal(responseData, &starGazerResponse)

	formattedOutput, _ := json.MarshalIndent(starGazerResponse, "", "\t")
	fmt.Println("\nRESULTS:")
	fmt.Printf("\n%s\n", string(formattedOutput))
}
