# Docker stuff

## Creating an AWS intance

https://docs.aws.amazon.com/AmazonECS/latest/developerguide/docker-basics.html

    sudo yum update -y
    sudo amazon-linux-extras install docker
    sudo service docker start
    sudo usermod -a -G docker ec2-user

re-login

    docker info

## Install golang and sources on host

[Doc](https://www.callicoder.com/docker-golang-image-container-example/)

    sudo yum install git
    # sudo yum install golang
    sudo yum remove golang
    cd ~
    wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go1.12.5.linux-amd64.tar.gz

    sudo vim /etc/profile

add
export PATH=$PATH:/usr/local/go/bin
export GOPATH=/home/ec2-user/go

re-login

    go get -d -v -t github.com/zew/go-questionnaire
    cd $GOPATH/src/github.com/zew/go-questionnaire
    cd /home/ec2-user/go/src/github.com/zew/go-questionnaire
    go build

Test

    ./go-questionnaire
    rm -rf ~/logdir-test
    mkdir -p ~/logdir-test
    export LOG_FILE_LOCATION=~/logdir-test/app.log
    echo ${LOG_FILE_LOCATION}
    ./go-questionnaire  >${LOG_FILE_LOCATION} 2>&1 &
    ps aux | grep go-questionnaire
    tail ${LOG_FILE_LOCATION} -n 40

open another shell

    wget http://localhost:8081/survey/generate-questionnaire-templates

## Create image, run container

    docker build -t dockered-qst .

    docker image ls

    # inspect image file system: stackoverflow.com/questions/44769315/
    docker run -it dockered-qst sh

### Simplest Dockerfile
    docker run -d  -p 8081:8081  dockered-qst

### Dockerfile with external app logging

    rm -  rf ~/log-go-quest-container
    mkdir -p ~/log-go-quest-container
    docker run -d  -p   80:8081  -v ~/log-go-quest-container:/root/logdir  dockered-qst

### Dockerfile with external logging and config

    rm   -rf ~/log-go-quest-container
    mkdir -p ~/log-go-quest-container
    rm   -rf ~/cfg-go-quest-container
    mkdir -p ~/cfg-go-quest-container

    cp ./config-example.json ~/cfg-go-quest-container/config.json
    cp ./logins-example.json ~/cfg-go-quest-container/logins.json

    docker run -d  -p   80:8081 \
       -v ~/log-go-quest-container:/root/logdir \
       -v ~/cfg-go-quest-container:/root/cfg    \
      dockered-qst

    docker container ls
    docker container stop [first 3 chars container ID]

### config and logins on host

SET CONFIG_FILE=.\cfg\config.json
SET LOGINS_FILE=.\cfg\logins.json

### Debug

    tail ~/log-go-quest-container/app.log -n 40

    docker container ls

    wget http://localhost:8081/survey/generate-questionnaire-templates

    sudo docker ps -a
    sudo docker logs [container ID first 3 chars]

### Error timezone

`Your location name must be valid ... unknown time zone Europe/Berlin`
There is no timezone file in alpine linux.

Fixed by using plplot/debian-latest-python34

