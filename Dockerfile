# Use an official Go runtime as a parent image
FROM golang:1.18

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Install any necessary dependencies
RUN go get -d -v ./...
RUN go install -v ./...

# Expose port 8888 for the API
EXPOSE 8888

# Run the API when the container starts
CMD ["go", "run", "app.go"]
