FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go get github.com/gin-gonic/gin
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go build -o main .

ENV DB_TYPE=local
ENV TYPE=operator

EXPOSE 8080
EXPOSE 27017

CMD ["/app/main"]