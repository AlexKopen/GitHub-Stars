# Github Stars Server

## Overview
The server for the GitHub Stars application, written in Go.  This application allows users
to pass in a list of "\<organization>/\<repository>" strings to a REST endpoint, receiving the total number
of stars for each repository in response.  Upon running, a HTTP server is launched on port 8080, allowing the user to 
make POST requests to `localhost:8080/api/v1/stars`.

[Gin](https://github.com/gin-gonic/gin) is used as the HTTP framework for building the API, and
[go-github](https://github.com/google/go-github) is used to interface with the GitHub API.

## Using with Docker compose
```
docker-compose up
```

## Building
```
docker build --target server-build -t github-stars-server .
```

## Running
Start the server at `localhost:8080`
```
docker run -p 8080:8080 -v github-stars-vol  github-stars-server
```

## Tests
```
docker build -t github-stars-server-test . && docker run github-stars-server-test
```

## Endpoints
```
POST /api/v1/stars
```

**Example Request Payload:**
```json
{
  "Repositories": [
    "angular/angular",
    "AlexKopen/AUR-Package-Scraper",
    "gin-gonic/examples",
    "Invalid",
    "doesnot/exist123"
  ]
}
```

**Example Response Payload:**
```json
{
    "TotalCount": 69955,
    "StarGazerResults": [
        {
            "Repository": "Invalid",
            "Count": 0,
            "ErrorMessage": "Invalid repository format.  Name must contain an owner and repo separated by a '/'."
        },
        {
            "Repository": "doesnot/exist123",
            "Count": 0,
            "ErrorMessage": "There was an error fetching this repository information from GitHub. Make sure the repository exists and the owner and repo name is correct."
        },
        {
            "Repository": "AlexKopen/AUR-Package-Scraper",
            "Count": 1,
            "ErrorMessage": ""
        },
        {
            "Repository": "gin-gonic/examples",
            "Count": 1131,
            "ErrorMessage": ""
        },
        {
            "Repository": "angular/angular",
            "Count": 68823,
            "ErrorMessage": ""
        }
    ]
}
```

## Error Handling
Making requests to any endpoint other than a POST to `/api/v1/stars` will result in a 404 response.  Sending an invalid JSON string in the request body will result in the following 400 response:
```json
{
    "Error": "Invalid JSON format in request."
}
```

## Data Models
```go
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
```

## Minikube
To run this Docker image locally using [minikube](https://minikube.sigs.k8s.io/docs/), follow these steps:
1. Set your local minikube environment to use images produced by Docker:
```
eval $(minikube docker-env)
```

2. Build the Docker image using the build instructions above.  After, run the image as a pod while exposing port `8080`:
```
kubectl run github-stars-pod --image=github-stars-server --port=8080 --image-pull-policy=Never
```

3. View the running pod:
```
kubectl get pods -o wide
```

or view the pod in a web application dashboard:
```
minikube dashboard
```
