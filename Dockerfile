FROM golang:latest

WORKDIR /go/src/app/

COPY . .

RUN go mod download -x

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 8080

ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main"
