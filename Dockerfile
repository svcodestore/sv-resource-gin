FROM golang:1.17 AS API
WORKDIR /sv-resource-gin
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server ./main.go && mkdir run && mv ./server ./config.yaml run

FROM alpine:latest
ENV GIN_MODE=release

COPY --from=API /sv-resource-gin/run ./
CMD ./server
