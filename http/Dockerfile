FROM alpine
LABEL author="ifanaticfire@gmail.com" date="2021-10-16"

WORKDIR . /bin

ADD bin/amd64/httpserver /httpserver
EXPOSE 80
ENTRYPOINT /httpserver