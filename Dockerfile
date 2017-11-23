FROM golang:1.9
MAINTAINER Alexander R. <spaizadv@gmail.com>
ARG REPOSITORY_PATH=${GOPATH}/src/github.com/panoplyio/cwlogs/
COPY . ${REPOSITORY_PATH}
WORKDIR ${REPOSITORY_PATH}
RUN go get -v ./... && go install
RUN mkdir -p /logs
WORKDIR /logs
ENTRYPOINT ["cwlogs"]
