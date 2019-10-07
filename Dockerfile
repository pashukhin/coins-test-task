FROM golang:latest
RUN mkdir -p /go/src/github.com/pashukhin/coins-test-task
ADD . /go/src/github.com/pashukhin/coins-test-task
WORKDIR /go/src/github.com/pashukhin/coins-test-task/cmd/test_task
RUN go get -v