FROM golang:alpine As builder

WORKDIR /app
ADD https://raw.githubusercontent.com/BwayCer/k8s.test/master/cheerAndPanicServer/src/main.go .
RUN go build -o app


FROM alpine

WORKDIR /var/www/app
COPY --from=builder /app .
CMD ./app

