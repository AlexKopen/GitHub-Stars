# Github Stars Client

## Overview
The client for the GitHub Stars application, written in Go.  This application allows users
to pass in a list of "\<organization>/\<repository>" strings via a cli, outputting the total number
of stars for each repository in response.

## Running
With the server running at `localhost:8080`, to start the client in a terminal:
```
cd pkg && go run .
```
### CLI Arguments
To run the client with error messages suppressed, set the `suppress` flag to true:
```
cd pkg && go run . -suppress=true
```

## Tests
With the server running at `localhost:8080`, to run the tests:
```
cd pkg && go test
```

## Error Handling
By default, the client will display any errors associated with invalid repository names to the user.

Ex.
```
Invalid repository name: asdf.  Name must contain exactly one '/'.  This input will not be processed.
```
