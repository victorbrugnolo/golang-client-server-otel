
FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o temp-by-zipcode-client ./cmd
ENTRYPOINT [ "./temp-by-zipcode-client" ]
