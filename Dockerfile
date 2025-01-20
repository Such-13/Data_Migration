# Start from Golang base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy all files into the container
COPY . .

# Install dependencies
RUN go mod tidy

# Build the application
RUN go build -o /app/main .

# Command to execute the container
CMD ["/app/main"]
