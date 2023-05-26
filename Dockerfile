FROM golang:1.20.2-alpine

RUN mkdir /apk
WORKDIR /app
ADD . /app


COPY  go.mod .
COPY go.sum .
COPY  .env .

# ENV GO111MODULE=on
# ENV GOFLAGS=-mod=vendor

RUN go mod download
RUN go build -o main .

EXPOSE 8080
CMD [ "/app/main" ]