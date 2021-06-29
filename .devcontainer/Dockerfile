#
#######################################################################################
#
# Development-Base-Image ...
FROM ubuntu:18.04
#
LABEL maintainer="Franz Branntvino"
LABEL Name=devbox
LABEL Version=0.0.1
#
#######################################################################################
# USER root for the apt-get sequences (based on ubuntu)
USER root
#
#######################################################################################
#
# as a user apt-get sequences can be called also with sudo apt-get OR use simply 'sudo -i' to become root (including the change of HOME=/root) ...
# to change an existing user in general in the shell would be for instance 'su - user' ...
#USER user
# because of Kubernetes and Jenkins-pipeline issues, to stay consistent - we simply are root for now ...
#
# Define home directory:
ENV HOME /home/user
# Define working directory (will also be the directory for carrying out subsequent docker commands):
WORKDIR $HOME
#
#######################################################################################
#
ENV DEBIAN_FRONTEND=noninteractive
#
#######################################################################################
#
# Preparing the Docker Base Image for the Go and Docker API Development Environment ...
# it is always helpful sorting packages alpha-numerically to keep the overview ;)
RUN \
    apt-get update && \
    apt-get -y upgrade && \
    # Configure apt, install packages and tools
    apt-get -y install --no-install-recommends apt-utils dialog 2>&1 \
    && \
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    apt-get -y install \
        bison build-essential bzip2 \
        ca-certificates cmake curl \
        default-jdk \
        flex \
        git gnupg-agent \
        iproute2 \
        lsb-release \
        net-tools \
        openssl \
        p7zip-full patch pkg-config procps \
        software-properties-common sudo \
        wget \
        xz-utils \
    && \
    # Clean up
    apt-get autoremove -y && \
    apt-get clean -y && \
    rm -rf /tmp/* /var/tmp/* && rm -rf /var/lib/apt/lists/*
#
# Install Docker & Docker-Compose
RUN \
    apt-get update && \
    apt-get -y upgrade && \
    apt-get -y install --no-install-recommends apt-utils dialog 2>&1 \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release \
    && \
    # Clean up
    apt-get autoremove -y && \
    apt-get clean -y && \
    rm -rf /tmp/* /var/tmp/* && rm -rf /var/lib/apt/lists/*
RUN \
ARCH=amd64 && \
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg \
    && \
    echo \
        "deb [arch="$ARCH" signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
        $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null \
    && \
    apt-get update && \
    apt-get -y install docker-ce docker-ce-cli containerd.io
RUN \
VERSION=1.29.2 && \
    curl -L "https://github.com/docker/compose/releases/download/"$VERSION"/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && \
    chmod +x /usr/local/bin/docker-compose && \
    ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
RUN \
    docker-compose --version
#
# Install Go
RUN \
ARCH=amd64 && \
OS=linux && \
VERSION=1.16.5 && \
    cd ~ && \
    wget https://dl.google.com/go/go$VERSION.$OS-$ARCH.tar.gz && \
    sha256sum go$VERSION.$OS-$ARCH.tar.gz && \
    tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz && \
    rm go$VERSION.$OS-$ARCH.tar.gz \
    && echo ""
#
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/home/user/go
#
RUN \
    go version
#
# Install Go tools
# it is always helpful sorting packages alpha-numerically to keep the overview ;)
RUN echo "Go tools like outline, linter and more ..." \
    && go get -u -v golang.org/x/lint/golint \
    && go get -u -v golang.org/x/tools/cmd/goimports \
    && go get -u -v golang.org/x/tools/cmd/gorename \
    && go get -u -v golang.org/x/tools/cmd/guru \
    && go get -u -v golang.org/x/tools/gopls
# gocode-gomod (https://github.com/stamblerre/gocode) an autocompletion daemon for go ...
RUN echo "gocode-gomod - an autocompletion daemon for go ..." \
    && go get -u -v github.com/stamblerre/gocode
RUN echo "..." \
    && go get -u -v github.com/acroca/go-symbols \
    && go get -u -v github.com/cweill/gotests/... \
    && go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct \
    && go get -u -v github.com/fatih/gomodifytags \
    && go get -u -v github.com/go-delve/delve/cmd/dlv
RUN echo "..." \
    && go get github.com/godoctor/godoctor@latest
RUN echo "..." \
    && go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint \
    && go get -u -v github.com/haya14busa/goplay/cmd/goplay \
    && go get -u -v github.com/josharian/impl \
    && go get -u -v github.com/mdempsky/gocode \
    && go get -u -v github.com/mgechev/revive \
    && go get -u -v github.com/ramya-rao-a/go-outline \
    && go get -u -v github.com/rogpeppe/godef \
    && go get -u -v github.com/sqs/goreturns
RUN echo "..." \
#    && go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
    && go get -u -v github.com/uudashr/gopkgs/v2/cmd/gopkgs
#    && go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
#    && go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs@latest
#    && go get github.com/uudashr/gopkgs/cmd/gopkgs@latest
RUN echo "..." \
    && go get -u -v github.com/zmb3/gogetdoc
RUN echo "..." \
    && go get -u -v honnef.co/go/tools/... \
    && echo ""
#
# Clean up
RUN echo "Clean up ..." \
    apt-get autoremove -y && \
    apt-get clean -y && \
    rm -rf /var/lib/apt/lists/* /go/src
#
#
#######################################################################################
#
# switch back to interactive ...
ENV DEBIAN_FRONTEND=teletype
#
SHELL ["/bin/bash", "-c"]
#
CMD ["bash"]
#