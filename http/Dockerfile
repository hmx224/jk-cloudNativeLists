FROM alpine
LABEL author="ifanaticfire@gmail.com" date="2021-10-16"

WORKDIR /app

ADD bin/amd64/httpserver /app/httpserver

EXPOSE 80

ENV NAME helloworld

ENTRYPOINT /app/httpserver