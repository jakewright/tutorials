FROM golang:latest
RUN go get -u golang.org/x/lint/golint
RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /go/src/github.com/jakewright/tutorials/home-automation/04-hue-lights/service.config
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD CompileDaemon -build="go install ." -command="/go/bin/service.config"