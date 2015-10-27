FROM golang

ADD . /go/src/github.com/frodebjerke/kafka-workshop
RUN go get github.com/Shopify/sarama
RUN go install github.com/frodebjerke/kafka-workshop
