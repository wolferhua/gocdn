FROM golang:1.9.4-stretch

ADD . /go


WORKDIR /go


RUN chmod +x start.sh

RUN  go build -v -o gocdn src/main/main.go

CMD ["/go/start.sh"]

#docker build -t wolferhua/gocdn  .


