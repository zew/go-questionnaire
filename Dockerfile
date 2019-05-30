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
# RUN go install -v ./...
# Instead:
# Build the Go app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/go-questionnaire .

RUN mkdir -p /go-questionnaire/stuff/responses
COPY ./static               /go-questionnaire/stuff/
COPY ./templates            /go-questionnaire/stuff/
# COPY ./server.key           /go-questionnaire/stuff/
# COPY ./server.pem           /go-questionnaire/stuff/
COPY ./config-example.json  /go-questionnaire/stuff/config.json
COPY ./logins-example.json  /go-questionnaire/stuff/logins.json


# New stage - from scratch - leightweight linux
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
# Alternative for testing
# FROM plplot/debian-latest-python34:latest
# RUN apt-get install ca-certificates

WORKDIR /root/

# Copy the pre-built binary from the previous stage
COPY --from=buildOS /go/bin/go-questionnaire .
COPY --from=buildOS /go-questionnaire/stuff/ .
# we cannot accomplish this; we can only COPY or ADD from context
# RUN   touch        ./templates/site.css

# Build Arg
# ARG LOG_DIR=/var/log/go-questionnaire
ARG LOG_DIR=logdir
RUN mkdir -p ${LOG_DIR}
ENV LOG_FILE_LOCATION=./${LOG_DIR}/app.log 

# a data volume is a directory in the container - breaking out of it
VOLUME ["./${LOG_DIR}"]

# This container exposes port 8081 to the outside world
EXPOSE 8081

# Run the binary program produced by `go build`
# CMD ["./go-questionnaire"]
CMD ["./go-questionnaire", ">${LOG_FILE_LOCATION} 2>&1"]
# CMD ["./go-questionnaire", ">${LOG_FILE_LOCATION} 2>&1 &"]