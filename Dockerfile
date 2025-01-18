FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Install dependencies and Air
RUN apk add --no-cache git bash && \
    go install github.com/air-verse/air@latest

# Copy go mod and sum files, then install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application files
COPY . .

# Expose the port your app is running on
EXPOSE 3000

# Command to run the app with Air
CMD ["air"]
