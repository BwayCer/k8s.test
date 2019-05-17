
FROM golang:alpine As builder

WORKDIR /app
COPY ./throb/golang/main.go .
RUN go build -o app


FROM alpine

WORKDIR /app
COPY --from=builder /app/app .

CMD ["./app"]

