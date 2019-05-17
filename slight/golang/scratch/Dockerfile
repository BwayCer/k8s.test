
FROM golang:alpine As builder

WORKDIR /app
COPY ./throb/golang/main.go .
RUN go build -o app


FROM scratch
COPY --from=builder /app/app /app
CMD ["/app"]

