FROM golang:1.21-alpine
LABEL authors="1"

RUN go version
ENV GOPATH=/

COPY .. ./

RUN go mod download
RUN go build -0 awesome./H_kitchen.go

CMD [".awesome"]
ENTRYPOINT ["top", "-b"]