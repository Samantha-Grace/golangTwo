FROM golang:1.11

WORKDIR /go/src/github.com/Samantha-Grace/golangTwo
COPY . .

RUN go install
EXPOSE 8080
CMD ["go-short"]