# template from flaviocopes.com/golang-docker/
FROM golang:1.12 as buildOS

# Add Maintainer Info
LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"

# Build Args
ARG APP_NAME=go-questionnaire
ARG LOG_DIR=/${APP_NAME}/logs

# Create log dir
RUN mkdir -p ${LOG_DIR}

# Environment vars
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log 

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

RUN mkdir -p /go-questionnaire/stuff
COPY ./static               /go-questionnaire/stuff/
COPY ./templates            /go-questionnaire/stuff/

# we cannot accomplish this; we can only COPY or ADD from context
# RUN   touch                 /go-questionnaire/stuff/templates/site.css

COPY ./config-example.json  /go-questionnaire/stuff/config.json
COPY ./logins-example.json  /go-questionnaire/stuff/logins.json



# New stage - from scratch - leightweight linux

# FROM plplot/debian-latest-python34:latest
# RUN apt-get --no-cache add ca-certificates

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy the pre-built binary from the previous stage
COPY --from=buildOS /go/bin/go-questionnaire .
COPY --from=buildOS /go-questionnaire/stuff/ .

# This container exposes port 8081 to the outside world
EXPOSE 8081

# Declare volumes to mount
# VOLUME ["/go-questionnaire-volume/logs"]

# Run the binary program produced by `go build`
CMD ["go-questionnaire"]