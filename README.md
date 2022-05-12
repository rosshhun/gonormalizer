# gonormalizer
[![Go Reference](https://pkg.go.dev/badge/github.com/rosshhun/gonormalizer.svg)](https://pkg.go.dev/github.com/rosshhun/gonormalizer)
[![Go Report Card](https://goreportcard.com/badge/github.com/rosshhun/gonormalizer)](https://goreportcard.com/report/github.com/rosshhun/gonormalizer)

A simple and easy to use URL Normalisation package for GO.

**Note:** This package does not do URL sanitization or Validation. It's important and recommended to do URL validation before passing url's as function parameters, the package only deals with normalization. If you are using this package in server context it's upto you to perform validation to prevent from various attacks. More information related to normalization can be found here:
> [Normalize](https://en.wikipedia.org/wiki/URL_normalization)\
> [RFC3986](https://www.ietf.org/rfc/rfc3986.txt)

## Install

```sh
go get github.com/rosshhun/gonormalizer
```

## Usage/Examples

```go
import (
	"github.com/rosshhun/gonormalizer"
)

func main(){
	u, err := gonormalizer.ForceHttps("http://example.com/")
	if err != nil{
		fmt.Prinltln(err)
	}
	fmt.Prinltln(u)
	//=>https://example.com/
}

```

## Function Signatures

### functionName(string) (string, error)

#### Parameter: `string`
Pass url as a string to function parameter, it's important to validate the url, with regex or other libraries Available.
#### Return Type: `string`
The value of String return type is a 'url' if there are no errors occured, if there is an error occured then the value of the string return type will be "".
#### Return Type: `error`
The value of Error return type is a 'nil' if there are no errors occured, if there is an error then the value of error return type  will be a custom erorr.


### functionName(string, string) (string, error)

#### Parameter1: `string`
Pass url as a string to function parameter, it's important to validate the url, with regex or other libraries Available.
#### Parameter2: `string`
Pass port number or protocol as a string to function parameter, it's important to validate the url, with regex or other libraries Available.
#### Return Type: `string`
The value of String return type is a 'url' if there are no errors occured, if there is an error occured then the value of the string return type will be "".
#### Return Type: `error`
The value of Error return type is a 'nil' if there are no errors occured, if there is an error then the value of error return type  will be a custom erorr.

### functionName(string) bool

#### Parameter: `string`
Pass url as a string to function parameter, it's important to validate the url, with regex or other libraries Available.
#### Return Type: `bool`
The value of bool return type can be a 'true/fale'

### functionName(string) string

#### Parameter: `string`
Pass url as a string to function parameter, it's important to validate the url, with regex or other libraries Available.
#### Return Type: `string`
The value of String return type is a 'modified url'.


## List of Functions
```
func AddPort(u string, p string) (string, error)
func AddProtocol(u string, p string) (string, error)
func AddTrailingDot(u string) (string, error)
func AddTrailingSlash(u string) (string, error)
func DefaultProtocol(u string) (string, error)
func ForceHttp(u string) (string, error)
func ForceHttps(u string) (string, error)
func IsEmpty(u string) bool
func IsValid(u string) bool
func LowerCase(u string) string
func RemoveTrailingDot(u string) (string, error)
func RemoveTrailingSlash(u string) (string, error)
func Scheme(u string) (string, error)
func StripAuthentication(u string) (string, error)
func StripHash(u string) (string, error)
func StripPort(u string) (string, error)
func StripProtocol(u string) (string, error)
func StripTextFragment(u string) (string, error)
func StripWWW(u string) (string, error)
func TrimURL(u string) string
```
