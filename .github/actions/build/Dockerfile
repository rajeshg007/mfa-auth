FROM golang:1.14.6-buster

LABEL name="Golang Build Action"
LABEL maintainer="Rajesh Guduguntla"
LABEL version="0.0.1"

LABEL com.github.actions.name="Golang Build Action"
LABEL com.github.actions.description="Run Golang commands"
LABEL com.github.actions.icon="box"
LABEL com.github.actions.color="blue"

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

CMD ""