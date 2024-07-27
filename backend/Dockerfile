# use the official Golang image as the base image
From golang:1.20 as builder

# set the current working directory inside the container
WORKDIR '/bugsfounder'

# copy the go modules manifests
COPY go.mod go.sum ./

# Download dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download 

# copy the source code into the container
COPY . .

# build the go app
RUN go build -0 main .

# Second stage: Create a small image for the app 
FROM alpine:latest 

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]