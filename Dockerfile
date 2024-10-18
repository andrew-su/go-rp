# Build container
FROM golang:latest AS build

COPY main.go main.go
RUN GOOS=linux GOARCH=amd64 go build -o /app main.go

# Final container
FROM gcr.io/distroless/base:nonroot
EXPOSE 8080
COPY --from=build /app /
CMD ["/app"]
