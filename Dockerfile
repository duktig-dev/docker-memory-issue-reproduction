FROM golang:1.17

WORKDIR /app
COPY ./src/go.mod ./
COPY ./src/go.sum ./
COPY ./src /app

RUN go mod download

# go build -o {output-bin-file} {source-files-path}
RUN go build -o /testing-app /app/cmd/

CMD ["/testing-app"]
