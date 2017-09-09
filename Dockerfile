FROM golang:1.9-alpine

RUN mkdir -p /go/src/github.com/MOZGIII/aws-ecr-login
WORKDIR /go/src/github.com/MOZGIII/aws-ecr-login

COPY . .
RUN go install ./cmd/aws-ecr-login

CMD ["aws-ecr-login"]
