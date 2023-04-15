FROM golang:1.20

WORKDIR /go/go-with-kafka

# Install librdkafka to use kafka in go code
RUN apt-get update && \
    apt-get install -y librdkafka-dev

# To run your container withput crash for tests
CMD [ "tail", "-f", "/dev/null" ]