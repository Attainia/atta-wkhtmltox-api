# ATTA WKHTMLTOX API

![](https://travis-ci.org/Attainia/atta-wkhtmltox-api.svg?branch=master)

WKHTMLTOX API is an http wrapper api round the [wkhtmltopdf](https://wkhtmltopdf.org/) cli tool.

## Build
```bash
docker build -t wkhtmltox-api .
```

## Run
```bash
docker run wkhtmltox-api -p 80:80 .
```

## Test
```bash
docker build -f Dockerfile.test -t wkhtmltox-api-test . \
    && docker run wkhtmltox-api-test go test -v wkhtmltox-api/cmd \
    && docker run wkhtmltox-api-test go test -v wkhtmltox-api/internal
```

## Usage

### HTML to PDF

Request
```http request
POST http://localhost:80
Content-Type: text/html
Accept: application/pdf

<html>
    <body>
        <h1>Hello World</h1>
    </body>
</html>
```

Response
```http request
HTTP/1.1 200 OK
Content-Type: application/pdf

<pdf data>
```
