# template from flaviocopes.com/golang-docker/
FROM golang:1.12 as buildOS

# Add Maintainer Info
LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"


# Set current working directory inside the container
WORKDIR $GOPATH/src/github.com/zew/go-questionnaire

# Copy everything from current host directory to the PWD (present working dir) inside the container
# WHY is this neccessary?
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8081 to the outside world
EXPOSE 8081


# Run the binary program produced by `go install`
CMD ["go-questionnaire"]