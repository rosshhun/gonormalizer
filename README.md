# gonormalizer
[![Go Reference](https://pkg.go.dev/badge/github.com/rosshhun/gonormalizer.svg)](https://pkg.go.dev/github.com/rosshhun/gonormalizer)
[![Go Report Card](https://goreportcard.com/badge/github.com/rosshhun/gonormalizer)](https://goreportcard.com/report/github.com/rosshhun/gonormalizer)
A simple and easy to use URL Normalisation package for GO.

**Note:** This package does not do URL sanitization or Validation. It's important and recommended to do URL validation before passing url's as function parameters, the package only deals with normalization. If you are using this package in server context it's upto you to perform validation to prevent from various attacks. More information related to normalization can be found here:
> [Normalize](https://en.wikipedia.org/wiki/URL_normalization)
> [RFC3986](https://www.ietf.org/rfc/rfc3986.txt)

## Install

```sh
go get github.com/rosshhun/gonormalizer
```

## List of Functions Available
```
AddPort
AddProtocol
AddTrailingDot
AddTrailingSlash
DefaultProtocol
ForceHttp
ForceHttps
IsEmpty
IsValid
LowerCase
RemoveTrailingDot
RemoveTrailingSlash
Scheme
StripAuthentication
StripHash
StripPort
StripProtocol
StripTextFragment
StripWWW
TrimURL

```

## Usage

```go
import (
	"github.com/rosshhun/gonormalizer"
)

func main(){
	fmt.Prinltln(gonormalizer.ForceHttps("http://example.com/"))
	//=>https://example.com/
}

```

## API

### functionName(url) (string, error)

#### URL Type: `string`
Pass url as a string to function parameter, it's important to validate the url, with regex or other libraries Available.
#### Return Type: `string`
The value of String return type is a 'url' if there are no errors occured, if there is an error occured then the value of the string return type will be "".
#### Error Type: `string`
The value of Error return type is a 'nil' if there are no errors occured, if there is an error then the value of error return type  will be a custom erorr.