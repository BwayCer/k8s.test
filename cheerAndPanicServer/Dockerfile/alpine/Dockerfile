FROM golang

WORKDIR /var/www/app
ADD https://raw.githubusercontent.com/BwayCer/k8s.test/master/cheerAndPanicServer/src/main.go .
RUN go build .
CMD ./app

