
FROM golang:1.18-buster as builder
WORKDIR /go/src/lambda
COPY go.mod go.sum ./
RUN go mod download
COPY ./ .
RUN go build -a -o /main .

FROM debian:buster as runner
WORKDIR /app/
COPY --from=builder /main /main
ENTRYPOINT [ "/main" ]

FROM runner as runner-with-rie
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/bin/aws-lambda-rie
RUN chmod 755 /usr/bin/aws-lambda-rie
ENTRYPOINT [ "/usr/bin/aws-lambda-rie" ]