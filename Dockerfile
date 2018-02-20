FROM golang:1.9.4-stretch

ADD . /gocdn


WORKDIR /gocdn


RUN chmod +x start.sh


CMD ["./start.sh"]
