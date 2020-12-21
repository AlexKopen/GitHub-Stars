package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBaseStarsEndpoint(t *testing.T) {
	// Set up the test server and defer its closing
	ts := httptest.NewServer(initializeServer())
	defer ts.Close()

	// Call the API with a payload known to be successful
	resp, err := http.Post(fmt.Sprintf("%s/stars", ts.URL), "application/json", bytes.NewBuffer([]byte(SuccessfulRequest)))

	// Verify there was no error calling the endpoint
	if err != nil {
		t.Fatalf("Expected a succecssful request, but received the following error: %v", err)
	}

	// Check to make sure the response status code is 200
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but received %v", resp.StatusCode)
	}

	// Verify the content-type header is properly set
	val, ok := resp.Header["Content-Type"]
	if !ok {
		t.Fatalf("Expected the 'Content-Type' header to be set")
	}

	// Verify the content-type is JSON
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected 'application/json; charset=utf-8', but received %s", val[0])
	}

	//	Ensure the total count is greater than 0
	responseData, _ := ioutil.ReadAll(resp.Body)
	starGazerResponse := StarGazerResponse{}
	_ = json.Unmarshal(responseData, &starGazerResponse)

	// There are no valid results, don't return the response
	if starGazerResponse.TotalCount <= 0 {
		t.Fatalf("Expected the count to be greater than 0, but instead received a count of: %d", starGazerResponse.TotalCount)
	}
}

func TestStarsEndpointInvalidPayload(t *testing.T) {
	ts := httptest.NewServer(initializeServer())
	defer ts.Close()

	// Call using an invalid payload
	resp, err := http.Post(fmt.Sprintf("%s/stars", ts.URL), "application/json", bytes.NewBuffer([]byte(InvalidPayloadRequest)))

	if err != nil {
		t.Errorf("Expected a succecssful request, but received the following error: %v", err)
	}

	// Expect a 400 to be returned after an invalid payload is sent in
	if resp.StatusCode != 400 {
		t.Errorf("Expected status code 400, but received %v", resp.StatusCode)
	}
}

func TestStarsEndpointRepositoryFormatErrorMessage(t *testing.T) {
	ts := httptest.NewServer(initializeServer())
	defer ts.Close()

	// Call using an invalid <organization>/<repository> format
	resp, err := http.Post(fmt.Sprintf("%s/stars", ts.URL), "application/json", bytes.NewBuffer([]byte(RepositoryFormatErrorMessageRequest)))

	if err != nil {
		t.Fatalf("Expected a succecssful request, but received the following error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, but received %v", resp.StatusCode)
	}

	// Extract the JSON data and check to make sure the correct error message was sent in the payload
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var data StarGazerResponse
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		t.Fatalf("There was an error processing the response payload")
	}

	if data.StarGazerResults[0].ErrorMessage != RepositoryFormatErrorMessage {
		t.Errorf("The wrong error message was produced.  Expected '%s' but received '%s'.", RepositoryFormatErrorMessage, data.StarGazerResults[0].ErrorMessage)
	}
}

func TestStarsEndpointNonExistentGitHubRepoErrorMessage(t *testing.T) {
	ts := httptest.NewServer(initializeServer())
	defer ts.Close()

	// Call using a non-existent repository
	resp, err := http.Post(fmt.Sprintf("%s/stars", ts.URL), "application/json", bytes.NewBuffer([]byte(NonExistentGithubRepoErrorMessagesRequest)))

	if err != nil {
		t.Fatalf("Expected a succecssful request, but received the following error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, but received %v", resp.StatusCode)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var data StarGazerResponse
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		t.Fatalf("There was an error processing the response payload")
	}

	if data.StarGazerResults[0].ErrorMessage != GitHubErrorMessage {
		t.Errorf("The wrong error message was produced.  Expected '%s' but received '%s'.", GitHubErrorMessage, data.StarGazerResults[0].ErrorMessage)
	}
}
