# Github Stars Client

## Overview
The client for the GitHub Stars application, written in Go.  This application allows users
to pass in a list of "\<organization>/\<repository>" strings via a terminal client, outputting the total number
of stars for each repository in response.

## Running
To start the client in a terminal:
```
go run .
```
### CLI Arguments
To run the client with error messages suppressed, set the `suppressed` flag to true:
```
go run . -suppress=true
```

## Tests
With the server running at `localhost:8080`, run the tests:
```
go test
```

## Error Handling
By default, the client will display any errors associated with invalid repository names to the user.

Ex.
```
Invalid repository name: asdf.  Name must contain exactly one '/'.  This input will not be processed.
```
