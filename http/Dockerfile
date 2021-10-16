FROM alpine
LABEL author="ifanaticfire@gmail.com" date="2021-10-16"

WORKDIR /bin

ADD bin/amd64/httpserver /bin/httpserver

EXPOSE 80

ENV NAME helloworld

ENTRYPOINT /httpserver