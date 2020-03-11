FROM golang:1.11-alpine3.8 as gobuild
WORKDIR /go/src/wkhtmltox-api
RUN apk add --no-cache curl git gcc musl-dev \
    && curl -s https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
COPY . .
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /build/wkhtmltox-api cmd/wkhtmltox-api.go

FROM ubuntu
MAINTAINER Attainia <developers@attainia.com>
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        xvfb \
        libfontconfig \
        wkhtmltopdf
COPY --from=gobuild /build/wkhtmltox-api /usr/local/bin/wkhtmltoxapi
RUN chmod +x /usr/local/bin/wkhtmltoxapi
VOLUME "/data"
CMD ["/usr/local/bin/wkhtmltoxapi"]