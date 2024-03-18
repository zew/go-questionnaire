FROM ubuntu:latest

RUN apt-get update && \
    apt-get install -y wget git

# install golang
# https://go.dev/doc/install
RUN wget -P /tmp https://go.dev/dl/go1.22.1.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf /tmp/go1.22.1.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
