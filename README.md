# lz-go/config

A strongly opinionated Go configuration management

## Installation

```shell
go get github.com/lz-go/config
```

## Getting started

### Prepare config file(s)

Prepare the default `config.yaml` file, e.g.

```yaml
app:
  env: local
  service_name: example

http:
  host: 127.0.0.1
  port: 8000
```

Optionally override using `.env` file, e.g.

```text
APP_ENV=testing
HTTP_HOST=0.0.0.0
```

More override example can be found lately within this document.

### Load the config

This package needs to be loaded before using, e.g.

```go
package main

import (
	"log"

	"github.com/lz-go/config"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	// ... rest of your application
}
```

### Use the config

Once loaded, the config can be accessed anywhere within your application, e.g.

```go
package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lz-go/config"
)

func Serve() {
	host := config.GetStringOr("http.host", "localhost")
	port := config.GetIntOr("http.port", 8000)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil); err != nil {
		log.Fatal(err)
	}
}
```

## Receipts

### Override the config using environment variables

You may want to override the config via environment variables, e.g.

```shell
export APP_ENV=testing
```

This is also applied to Kubernetes deployment, e.g.

```yaml
apiVersion: apps/v1
kind: Deployment
spec:
  selector: {}
  template:
    spec:
      containers:
        - name: application
          env:
            - name: APP_ENV
              value: testing
```
