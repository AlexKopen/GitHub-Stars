FROM golang:1.15.6

WORKDIR /github-stars

COPY . .

RUN go get -d -v ./...

RUN go build -o ./out/github-stars .

EXPOSE 8080

CMD ["./out/github-stars"]