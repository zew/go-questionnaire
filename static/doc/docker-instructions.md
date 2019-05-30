# Docker stuff

## Creating an AWS intance

https://docs.aws.amazon.com/AmazonECS/latest/developerguide/docker-basics.html

    sudo yum update -y
    sudo amazon-linux-extras install docker
    sudo service docker start
    sudo usermod -a -G docker ec2-user

re-login

    docker info
    docker build -t go-questionnaire .

## Install golang and sources on host

[Doc](https://www.callicoder.com/docker-golang-image-container-example/)

    sudo mkdir -p $GOPATH/src/github.com/zew/go-questionnaire
    sudo yum install git
    # sudo yum install golang
    sudo yum remove golang
    cd ~
    wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go1.12.5.linux-amd64.tar.gz

    sudo vim /etc/profile

add  `export PATH=$PATH:/usr/local/go/bin`

    go get -d -v -t github.com/zew/go-questionnaire
    cd $GOPATH/src/github.com/zew/go-questionnaire
    cd /home/ec2-user/go/src/github.com/zew/go-questionnaire
    go build
    ./go-questionnaire

open another shell

    wget http://localhost:8081/survey/generate-questionnaire-templates

## Create image, run container

    cd /home/ec2-user/go/src/github.com/zew/go-questionnaire

    docker build -t dockered-qst .

    docker image ls

    # either
    docker run -d  -p 8081:8081  dockered-qst

    # or
    sudo mkdir -p /var/log/go-questionnaire
    sudo touch /var/log/go-questionnaire/app.log
    sudo chown ec2-user:ec2-user /var/log/go-questionnaire/app.log
    docker run -d  -p 8081:8081  -v /var/log/go-questionnaire:/root/logdir  dockered-qst

    #or
    mkdir -p ~/log
    docker run -d  -p 8081:8081  -v ~/log:/root/logdir  dockered-qst

    docker container ls

    sudo journalctl -fu docker.service
    docker exec -it <mycontainer> bash

    wget http://localhost:8081/survey/generate-questionnaire-templates

    docker container stop fff93d13a484
