FROM alpine:latest

MAINTAINER Carlos Rivero <carlosriveroaro7@gmail.com>

WORKDIR "/opt"

ADD .docker_build/webpage-gopherjs /opt/bin/webpage-gopherjs
#ADD ./templates /opt/templates
#ADD ./static /opt/static

CMD ["/opt/bin/webpage-gopherjs"]

