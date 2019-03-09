# template from flaviocopes.com/golang-docker/
FROM golang:1.8.3 as builder
WORKDIR /go/src/github.com/zew/go-questionnaire
RUN go get -d -v golang.org/x/crypto/acme/autocert
RUN go get -d -v go.etcd.io/bbolt
RUN go get -d -v github.com/alexedwards/scs
RUN go get -d -v github.com/alexedwards/scs/boltstore
RUN go get -d -v github.com/alexedwards/scs/memstore
COPY findlinks.go  .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o findlinks .

FROM alpine:latest
# apk is the yum / apt-get of alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/zew/go-questionnaire/go-questionnaire .
CMD ["./go-questionnaire"]