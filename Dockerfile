# template from flaviocopes.com/golang-docker/
FROM golang:1.12 as buildOS

# Add Maintainer Info
LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"

# Set current working directory inside the container
WORKDIR $GOPATH/src/github.com/zew/go-questionnaire

# Copy everything from current host directory to the PWD (present working dir) inside the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Install the package
# RUN go install -v ./...
# Instead:
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/go-questionnaire .

RUN  mkdir -p /go-questionnaire/stuff/responses
COPY ./static               /go-questionnaire/stuff/static/
COPY ./templates            /go-questionnaire/stuff/templates/

# defaults for config and logins - 
# but it is hopeless, since /cfg will be emptied when becoming a volume - see below
COPY ./config-example.json  /go-questionnaire/stuff/cfg/config.json
COPY ./logins-example.json  /go-questionnaire/stuff/cfg/logins.json


# New stage - from scratch - leightweight linux
# FROM alpine:latest  
# RUN apk --no-cache add ca-certificates

# Alternative; avoding error 
#   Your location name must be valid ... unknown time zone Europe/Berlin
FROM plplot/debian-latest-python34:latest
RUN apt-get install ca-certificates

WORKDIR /root/

# Copy the pre-built binary from the previous stage
COPY --from=buildOS /go/bin/go-questionnaire .
COPY --from=buildOS /go-questionnaire/stuff  .

# we cannot accomplish touch; we can only COPY or ADD from context; still needed by the app?
# RUN   touch        ./templates/site.css

# Build Arguments
ARG LOG_DIR=logdir

RUN mkdir -p ${LOG_DIR}
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log 

# a data volume is a directory in the container - breaking out of container into host
VOLUME ["./${LOG_DIR}"]

# application takes config and logins from these environment variables
RUN mkdir -p cfg
ENV CONFIG_FILE=cfg/config.json
ENV LOGINS_FILE=cfg/logins.json

# a data volume is a directory in the container - breaking out of container into host
# previous contents of ./cfg will be destroyed
VOLUME ["./cfg"]

# This container exposes port 8081 to the outside world
EXPOSE 8081

# Run the binary program produced by `go build`
# if we want output redirection, we must prefix /bin/sh
# dont use trailing & - otherwise the container exits instantly
CMD ["/bin/sh", "-c", "./go-questionnaire > ${LOG_FILE_LOCATION} 2>&1"]
