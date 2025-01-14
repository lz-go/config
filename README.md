# lz-go/config

A strongly opinionated configuration management library for Go applications, simplifying access to configuration values from multiple sources.  This library prioritizes a hierarchical configuration approach using YAML files, environment variables, and `.env` files, making it easy to manage settings for different environments and deployments.

## Features

* **Hierarchical Configuration:** Uses YAML for structured configuration, allowing for nested values and logical organization.
* **Environment Variable Override:**  Easily override YAML settings with environment variables, providing flexibility for different deployments.
* **.env File Support:** Supports loading a `.env` file for local development and easy environment variable management.
* **Automatic Environment Variable Mapping:** Automatically maps environment variables to configuration keys, simplifying the override process.
* **Type-Safe Getters:** Provides type-safe getter functions for various data types (string, bool, int, float, duration, etc.), reducing the risk of type-related errors.
* **Default Values:** Offers getter functions with default values, ensuring your application behaves correctly even if a configuration key is missing.
* **Testable Design:**  Easy to reset and override configuration values during testing, facilitating isolated and predictable test cases.

## Installation

```shell
go get github.com/lz-go/config
```

## Usage

### 1. Configuration File (config.yaml)

```yaml
app:
  env: local
  service_name: example_service

http:
  host: 0.0.0.0
  port: 8000
```

### 2. Environment Variable Overrides

Override configuration values using environment variables. The library automatically maps environment variables to configuration keys by replacing dots (.) with underscores (_) and converting to uppercase. For example:

```shell
export APP_ENV=production
export HTTP_PORT=3000
```

### 3. .env File (Optional)

For local development, you can use a .env file in the same directory as your config.yaml to manage environment variables. The library will automatically load this file.

```text
APP_ENV=production
HTTP_PORT=3000
```

### 4. Loading the Configuration

In your main.go (or equivalent entry point), load the configuration using config.Load(). This function searches for config.yaml in the current working directory and its parent directories until it finds one. It also loads any .env file present in the same directory as the config.yaml.

```go
package main

import "github.com/lz-go/config"

func main() {
	config.Load()
	// ... rest of your application
}
```

### 5. Accessing Configuration Values

Use the provided getter functions to access configuration values anywhere in your application:

```go
package mypackage

import "github.com/lz-go/config"

func MyFunction() {
	env := config.GetString("app.env")
	port := config.GetIntOr("http.port", 8000) // Get with default value
	// ... use the configuration values
}
```

## Unit Testing

Reset viper before each test to ensure a clean configuration state. You can then set specific configuration values for testing purposes.

```go
package mypackage

import (
	"testing"

	"github.com/lz-go/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestMyFunction(t *testing.T) {
	viper.Reset() // Crucial for isolated tests
	viper.Set("app.env", "test")
	viper.Set("http.port", 9000)
	
	// Call your function that uses the configuration
    // ... make assertions based on expected behavior with test configuration
}
```

## Error Handling

The config.Load() function will log an error if the .env file loading fails, but it will not stop the application. However, if config.yaml fails to load, the application will terminate with a fatal error. This ensures that your application fails fast if essential configuration is missing. Consider providing fallback mechanisms if this behavior is undesirable in your application's context.

## Best Practices

* **Structured Keys:** Use a dot-separated notation for hierarchical configuration keys (e.g., app.env, http.port).

* **Clear Documentation:** Document your configuration options in the config.yaml file with comments.

* **Environment-Specific Settings:** Use environment variables to manage settings that vary between environments (e.g., database credentials, API keys).
