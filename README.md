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
func Normalize(s string) (string, error)
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

## Examples


#### AddPort

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.AddPort("http://example.com/", "80")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### AddProtocol

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.AddProtocol("example.com", "https")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### AddTrailingDot

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.AddTrailingDot("example.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### AddTrailingSlash

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.AddTrailingSlash("https://example.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### DefaultProtocol

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.DefaultProtocol("example.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### ForceHttp

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.ForceHttp("https://example.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### ForceHttps

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.ForceHttps("http://example.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### IsEmpty

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u := gonormalizer.IsEmpty("https://example.com")
	fmt.Println(u)
}

```

#### IsValid

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u := gonormalizer.IsValid("https://example.com")
	fmt.Println(u)
}
```
#### LowerCase

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u := gonormalizer.LowerCase("HtTPs://example.com")
	fmt.Println(u)
}

```

#### Normalize

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.Normalize("//example.com")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### StripTrailingDot

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.StripTrailingDot("https://example.com.")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### StripTrailingSlash

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.StripTrailingSlash("https://example.com/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### Scheme

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.Scheme("https://example.com/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### StripAuthentication

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.StripAuthentication("user:password@example.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
```

#### StripHash

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.StripHash("https://example.com#")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
```

#### StripPort

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.StripPort("https://example.com:80")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### StripProtocol

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.StripProtocol("https://example.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### StripTextFragment

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.StripTextFragment("https://example.com#*~text339-+")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### StripWWW

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u, err := gonormalizer.StripWWW("https://www.example.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

```

#### TrimURL

```go
package main

import (
	"fmt"

	"github.com/rosshhun/gonormalizer"
)

func main() {
	u := gonormalizer.TrimURL("    https://www.example.com    ")
	fmt.Println(u)
}
```

#### Notes
Documentation is available here: [pkg.go.dev](https://pkg.go.dev/github.com/rosshhun/gonormalizer).

#### Support
If you do have a contribution to the package, feel free to create a Pull Request or an Issue.

#### What to contribute
If you don't know what to do, there are some features and functions that need to be done

- [ ]  Refactor code
- [ ]  Edit docs and [README](https://github.com/rosshhun/gonormalizer/README.md): spellcheck, grammar and typo check
- [ ]  Implement benchmarking
- [ ]  Implement batch of examples
- [ ]  Look at forks for new features and fixes

#### Advice
Feel free to create what you want, but keep in mind when you implement new features:
- Code must be clear and readable, names of variables/constants clearly describes what they are doing
- Public functions must be documented and described in source file and added to README.md to the list of available functions
- There are must be unit-tests for any new functions and improvements
