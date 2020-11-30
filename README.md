# GoHooks

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#networking)  
[![PkgGoDev](https://pkg.go.dev/badge/github.com/averageflow/gohooks/gohooks)](https://pkg.go.dev/github.com/averageflow/gohooks/gohooks)
[![Build](https://img.shields.io/github/workflow/status/averageflow/gohooks/Test)](#)
[![Size](https://img.shields.io/github/languages/code-size/averageflow/gohooks)](#)
[![Maintainability](https://api.codeclimate.com/v1/badges/fa3f27e42986b329c2b2/maintainability)](https://codeclimate.com/github/averageflow/gohooks/maintainability)
[![codecov](https://codecov.io/gh/averageflow/gohooks/branch/master/graph/badge.svg?token=DK72X8ROZN)](https://codecov.io/gh/averageflow/gohooks)
[![Go Report Card](https://goreportcard.com/badge/github.com/averageflow/gohooks)](https://goreportcard.com/report/github.com/averageflow/gohooks)
[![License](https://img.shields.io/github/license/averageflow/gohooks.svg)](https://github.com/averageflow/gohooks/blob/master/LICENSE.md)
[![Issues](https://img.shields.io/github/issues/averageflow/gohooks)](#)

GoHooks make it easy to send and consume secured web-hooks from a Go application. A SHA-256 signature is created with the sent data plus an encryption salt and serves to validate on receiving, effectively making your applications only accept communication from a trusted party.


## Installation

Add `github.com/averageflow/gohooks/v2` to your `go.mod` file and then import it into where you want to be using the package by using:

```go
import (
    "github.com/averageflow/gohooks/v2/gohooks"
)
```


## Usage

Here I will list the most basic usage for the GoHooks. If you desire more customization please read the section below for more options.

#### Sending

The most basic usage for sending is:

```go
// Data can be any type, accepts interface{}
data := []int{1, 2, 3, 4} 
// String sent in the GoHook that helps identify actions to take with data
resource := "int-list-example"
// Secret string that should be common to sender and receiver
// in order to validate the GoHook signature
saltSecret := "0014716e-392c-4120-609e-555e295faff5"

hook := &gohooks.GoHook{}
hook.Create(data, resource, salt)

// Will return *http.Response and error
resp, err := hook.Send("www.example.com/hooks")
```

#### Receiving

The most basic usage for receiving is:

```go
type MyWebhook struct {
    Resource string `json:"resource"`
    Data []int `json:"data"`
}

var request MyWebhook
// Assuming you use Gin Gonic, otherwise unmarshall JSON yourself.
_ = c.ShouldBindJSON(&request)

// Shared secret with sender
saltSecret := "0014716e-392c-4120-609e-555e295faff5"
// Assuming you use Gin Gonic, obtain signature header value
receivedSignature := c.GetHeader(gohooks.DefaultSignatureHeader)

// Verify validity of GoHook
isValid := gohooks.IsGoHookValid(request, receivedSignature, saltSecret)
// Decide what to do if GoHook is valid or not.
```

## Customization

GoHooks use the custom header `X-GoHooks-Verification` to send the encrypted SHA string. You can customize this header by initializing the GoHook struct with the custom option `SignatureHeader`. 

Example: 
```go
hook := &gohooks.GoHook{ SignatureHeader: "X-Example-Custom-Header" }
```

GoHooks are by default not verifying the receiver's SSL certificate validity. If you desire this behaviour then enable it by initializing the GoHook struct with the custom option `IsSecure`.

Example: 
```go
hook := &gohooks.GoHook{ IsSecure: true }
```

GoHooks will by default be sent via a `POST` request. If you desire to use a different HTTP method, amongst the allowed `POST`, `PUT`, `PATCH`, `DELETE`, then feel free to pass that option when initializing the GoHook struct, with `PreferredMethod`. Any other value will make the GoHook default to a `POST` request.

Example: 
```go
hook := &gohooks.GoHook{ PreferredMethod: http.MethodDelete }
```

GoHooks can be sent with additional HTTP headers. If you desire this then initialize the GoHook struct with the custom option `AdditionalHeaders`.

Example: 
```go
hook := &gohooks.GoHook{ AdditionalHeaders: map[string]string{"X-Header-Test": "Header value"} }
```

If you want to send your payload without GoHooks modifying it into its struct, use `CreateWithoutWrapper` when creating the GoHook.