FROM golang:1.10

WORKDIR /go/src/github.com/Kiraco/don_gorogoto

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "get", "github.com/gorilla/mux"]
RUN ["go", "get", "github.com/xeipuuv/gojsonschema"]
RUN ["go", "get", "github.com/google/uuid"]



ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/api/" -command="./api"