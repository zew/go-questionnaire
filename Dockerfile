FROM --platform=linux/amd64 ubuntu:latest

RUN apt-get update && \
    apt-get install -y wget git build-essential

# install golang
# https://go.dev/doc/install
RUN wget -P /tmp https://go.dev/dl/go1.22.1.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf /tmp/go1.22.1.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

RUN mkdir /app
WORKDIR /app

COPY . .

RUN chmod 777 ./rebuild.sh
RUN ./rebuild.sh

EXPOSE 8081

CMD ["/bin/sh", "-c", "PORT=8080 ./go-questionnaire"]