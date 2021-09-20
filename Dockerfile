FROM golang:1.16-alpine

WORKDIR /app
COPY . /app

RUN go build -o /bin/test-task

EXPOSE 8080

ENTRYPOINT ["/bin/test-task"]
