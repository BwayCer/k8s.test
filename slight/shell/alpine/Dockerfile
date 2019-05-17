
FROM alpine

RUN apk add bash && \
    rm -rf /var/cache/apk/*

WORKDIR /app
COPY ./throb/shell/main.sh .

CMD ["bash", "./main.sh"]

