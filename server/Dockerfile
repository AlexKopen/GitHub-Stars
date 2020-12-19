FROM golang:1.15.6 AS server-build

VOLUME /github-stars-vol
WORKDIR /github-stars

COPY . .

RUN go get -d -v ./...
RUN go build -o ./out/github-stars .

EXPOSE 8080

CMD ["./out/github-stars"]

FROM golang:1.15.6

COPY --from=server-build /github-stars .
COPY --from=server-build /$GOPATH /$GOPATH

CMD ["go", "test"]
