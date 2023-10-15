FROM golang:alpine3.18

# Initiate a folder for the project called app
WORKDIR /app

# Move the files that are responsible for installing the dependencies 
COPY go.mod go.sum ./

# Install the dependencies
RUN go mod download

# Move the rest of the files
COPY . .

# Build the project into an executable
RUN go build -o main .

# Run the executable
CMD ["/app/main"]