FROM alpine:latest

MAINTAINER Holger Koser <holger.koser@gmail.com>

WORKDIR "/opt"

ADD .docker_build/code4k8s /opt/bin/code4k8s
ADD ./templates /opt/templates

CMD ["/opt/bin/code4k8s"]

