#Start Debian image with latest Go version
#Workspace at /go
FROM golang

ADD . /go/src/github.com/9uuso/vertigo

RUN go get github.com/tools/godep && go get -u github.com/9uuso/vertigo && cd /go/src/github.com/9uuso/vertigo && godep go build

WORKDIR /go/src/github.com/9uuso/vertigo

ENTRYPOINT PORT="80" MARTINI_ENV="production" vertigo

EXPOSE 80
