FROM golang:1.23-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -o /app/build/server ./cmd/server.go

FROM alpine

COPY --from=build /app/build/server /server

EXPOSE 8080

ENTRYPOINT [ "/server" ]
