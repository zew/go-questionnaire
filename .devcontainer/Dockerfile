FROM --platform=linux/amd64 ubuntu:latest

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
  apt-get install -y wget git build-essential zsh

# install golang
# https://go.dev/doc/install
RUN wget -P /tmp https://go.dev/dl/go1.22.1.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf /tmp/go1.22.1.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

# install oh-my-zsh
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -q -O - | zsh \
  && cp ~/.oh-my-zsh/templates/zshrc.zsh-template ~/.zshrc \
  && sed -i "/^plugins=/c\plugins=(git dotenv)" ~/.zshrc \
  #
  # link make to devcontainer makefile
  && echo 'alias make="make -f /workspace/.devcontainer/Makefile"' >> ~/.zshrc

ENV DEBIAN_FRONTEND=dialog