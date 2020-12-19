FROM golang:1.15.6 AS server-build

VOLUME /github-stars-vol
WORKDIR $GOPATH/github-stars

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY pkg pkg
WORKDIR $GOPATH/github-stars/pkg
RUN go build -o ./out/github-stars .

EXPOSE 8080

CMD ["./out/github-stars"]

FROM golang:1.15.6

COPY --from=server-build /$GOPATH /$GOPATH
WORKDIR $GOPATH/github-stars/pkg

CMD ["go", "test"]
