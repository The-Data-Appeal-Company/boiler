# Boiler 

[![Go Report Card](https://goreportcard.com/badge/github.com/The-Data-Appeal-Company/boiler)](https://goreportcard.com/report/github.com/The-Data-Appeal-Company/boiler)
![Go](https://github.com/The-Data-Appeal-Company/boiler/workflows/Go/badge.svg)
[![license](https://img.shields.io/github/license/The-Data-Appeal-Company/boiler.svg)](LICENSE)


Highly configurable cache warmer
## Usage 

```bash
boiler --config config.yml 
```

example configuration 
```yaml
source:
  type: database
  params:
    uri: 'sslmode=require user=<user> password=<password> host=<host> port=<port> dbname=<db>'
    driver: postgres
    url_column: 'uri'
    http_method_column: 'method'
    query: |-
       select uri, request_date, 'GET' as method FROM api_calls

transformations:
  - type: rewrite-host
    params:
      host: localhost:8080

  - type: write-header
    params:
      headers:
        X-User: Cache-Warmer-Boiler

executor:
  type: http
  params:
    continue_on_error: true
    concurrency: 1
    timeout: 60s

```

