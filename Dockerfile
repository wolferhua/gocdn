FROM golang:1.9.4-stretch

ADD . /go


WORKDIR /go


RUN chmod +x start.sh


CMD ["./start.sh"]
