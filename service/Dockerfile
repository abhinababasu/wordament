
FROM golang:latest

MAINTAINER abhinaba@gmail.com

COPY service /
ADD english0.dict /
ADD english2.dict /
WORKDIR /
CMD ["/service", "-port=8090"]
