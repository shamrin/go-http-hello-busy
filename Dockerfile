# Start with a minimal base image containing Go
FROM golang:1.23.3-alpine

WORKDIR /app
COPY go.mod hello.go ./
RUN GOOS=linux GOARCH=amd64 go build -o server .

# Expose the port that the application listens on
EXPOSE 80

# Command to run the application
CMD ["./server"]
