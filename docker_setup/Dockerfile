############################################################
# Dockerfile to build sandbox for executing user code
# Based on Ubuntu
############################################################

FROM ubuntu:18.04

# Update the repository sources list
RUN echo "deb http://archive.ubuntu.com/ubuntu bionic main universe" > /etc/apt/sources.list
RUN apt-get update
RUN apt update
#Install all the languages/compilers we are supporting.
# Install OpenJDK-8
RUN apt-get update && \
  apt-get install -y openjdk-8-jdk && \
  apt-get install -y ant && \
  apt-get clean;

# Fix certificate issues
RUN apt-get update && \
  apt-get install ca-certificates-java && \
  apt-get clean && \
  update-ca-certificates -f;

# Setup JAVA_HOME -- useful for docker commandline
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64/
RUN export JAVA_HOME

RUN apt-get install -y php
RUN apt-get install -y ruby
RUN apt-get install -y python3
RUN apt-get install -y python
RUN apt-get install -y scala
RUN apt-get install -y sudo
RUN apt-get install -y golang
RUN apt-get install -y clojure
RUN apt-get install -y bc
ENV NODE_VERSION=12.6.0
RUN apt install -y curl
RUN curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.34.0/install.sh | bash
ENV NVM_DIR=/root/.nvm
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"
RUN node --version
RUN npm config set strict-ssl false
RUN npm config set user 0
RUN npm config set unsafe-perm true
RUN npm install -g underscore request express jade shelljs passport http sys jquery lodash async mocha moment connect validator restify ejs ws co when helmet wrench brain mustache should backbone forever  debug && export NODE_PATH=/usr/local/lib/node_modules/
