# Build stage
FROM golang:1.22 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app main.go

# Run stage (lightweigth)
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/app .
EXPOSE 8080
CMD ["./app"]
