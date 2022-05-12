package gonormalizer

import (
	"errors"
	"net/url"
	"strings"
)

// AddPort attaches the specified port to the end of URL
// Accepts URL and Port Number as a string argument
// if string matches the patterns (pattern is regular expression) then
// returns the modified string with port number and nil
// In case of error return is empty string with a customized error
func AddPort(u string, p string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	if rxSlash.Match([]byte(u)) {
		u, _ = StripTrailingSlash(u)
	}
	if !rxPort.Match([]byte(u)) {
		return u + ":" + p, nil
	}
	return "", errors.New("Port already exist")
}

// AddProtocol attaches the specified Protocol to the URL
// Accepts URL and Protocol as a string argument
// if string matches the patterns (pattern is regular expression) then
// returns the modified string with protocol and nil
// In case of error return is empty string with a customized error
func AddProtocol(u string, p string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	if !rxHttp.Match([]byte(u)) && rxWWW.Match([]byte(u)) || !rxHttp.Match([]byte(u)) {
		return p + "://" + u, nil
	}
	return "", errors.New("Protocol already Exists")
}

// AddTrailingSlash attaches the / to the end of URL
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// returns the modified string and nil
// In case of error return is empty string with a customized error
func AddTrailingSlash(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	if !rxSlash.Match([]byte(u)) {
		return u + "/", nil
	}
	return "", errors.New("TrailingSlash Exist in URL")
}

// AddTrailingDot attaches the dot (.) to the end of URL
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// returns the modified string and nil
// In case of error return is empty string with a customized error
func AddTrailingDot(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	if rxSlash.Match([]byte(u)) {
		u, _ = StripTrailingSlash(u)
	}
	return u + ".", nil
}

// DefaultProtocol attaches the http:// to the URL
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// returns the modified string and nil
// In case of error return is empty string with a customized error
func DefaultProtocol(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	if !rxHttp.Match([]byte(u)) && rxWWW.Match([]byte(u)) || !rxHttp.Match([]byte(u)) {
		return "http://" + u, nil
	}
	return "", errors.New("Protocol already Exists")
}

// ForceHttp converts the URL from https to http
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// returns the modified string and nil
// In case of error return is empty string with a customized error
func ForceHttp(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	strTemp := rxFhttp.FindString(u)
	if rxFhttp.Match([]byte(u)) {
		return strings.Replace(u, strTemp, "http:", -1), nil
	}
	return "", errors.New("Protocol does not exist")
}

// ForceHttps converts the URL from http to https
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// returns the modified string and nil
// In case of error return is empty string with a customized error
func ForceHttps(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	strTemp := rxFhttps.FindString(u)
	if rxFhttps.Match([]byte(u)) {
		return strings.Replace(u, strTemp, "https:", -1), nil
	}
	return "", errors.New("Protocol does not exist")
}

// IsValid checks if URL is in format with the URL Pattern
// IsValid is used by every function in the library
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// returns true, else false
func IsValid(u string) bool {
	return rxC.Match([]byte(u))
}

// IsEmpty checks URL is empty or not
// IsEmpty is used by every function in the library
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// returns true, else false
func IsEmpty(u string) bool {
	u = strings.TrimSpace(u)
	return len(u) == 0
}

// LowerCase checks if the passed URL is in lowercase
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// it returns the lowercase string
func LowerCase(u string) string {
	return strings.ToLower(u)
}

// StripTrailingSlash removes TrailingSlash / from the end of URL
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// Returns the modified string and nil
// In case of error return is empty string with a customized error
func StripTrailingSlash(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	if rxSlash.Match([]byte(u)) {
		return u[:len(u)-1], nil
	}
	return "", errors.New("No TrailingSlash Exist")
}

// StripTrailingDot removes the dot (.) from the end of URL
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// Returns the modified string and nil
// In case of error return is empty string with a customized error
func StripTrailingDot(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	if rxTd.Match([]byte(u)) {
		return u[:len(u)-1], nil
	}
	return "", errors.New("TrailingDot does not exist")
}

// Scheme presents us with the scheme or portocol of the URL
// internally Scheme uses url.Parse and scheme functions from url package
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// Returns the modified string and nil
// In case of error return is empty string with a customized error
func Scheme(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	urlTemp, err := url.Parse(u)
	if err != nil {
		return "", errors.New("Not a vaild URL")
	}
	return urlTemp.Scheme, nil
}

// StripProtocol removes the protocol from URL
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// Returns the modified string and nil
// In case of error return is empty string with a customized error
func StripProtocol(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	strTemp := rxHttp.FindString(u)
	if rxHttp.Match([]byte(u)) {
		return strings.Replace(u, strTemp, "", -1), nil
	}
	return "", errors.New("No Protocol Exist")
}

// StripWWW removes the www. from URL
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// Returns the modified string and nil
// In case of error return is empty string with a customized error
func StripWWW(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	strTemp := rxWWW.FindString(u)
	if rxWWW.Match([]byte(u)) {
		return strings.Replace(u, strTemp, "", -1), nil
	}
	return "", errors.New("www does not exist")
}

// StripHash removes the # and contents after #example from URL
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// Returns the modified string and nil
// In case of error return is empty string with a customized error
func StripHash(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	strTemp := rxShash.FindString(u)
	if rxShash.Match([]byte(u)) {
		return strings.Replace(u, strTemp, "", -1), nil
	}
	return "", errors.New("Hash does not exist")
}

// StripTextFragment removes text fragments from the end of URL
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// Returns the modified string and nil
// In case of error return is empty string with a customized error
func StripTextFragment(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	strTemp := rxShash.FindString(u)
	if rxShash.Match([]byte(u)) {
		return strings.Replace(u, strTemp, "", -1), nil
	}
	return "", errors.New("Text Fragment does not exist")
}

// StripAuthentication removes authentication from the end of URL
// Expected input format
// "user:password@@example.com", "https://user:password@@example.com"
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// Returns the modified string and nil
// In case of error return is empty string with a customized error
func StripAuthentication(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	strTemp := rxSauth.FindString(u)
	strURL := rxHttp.FindString(u)
	if rxSauth.Match([]byte(u)) {
		x := strings.Replace(u, strTemp, "", -1)
		x = strURL + x
		return x, nil
	}
	return "", errors.New("Authentication does not exist")
}

// StripPort detaches the port from URL
// Accepts URL as a string argument
// if string matches the patterns (pattern is regular expression) then
// returns the modified string with port number and nil
// In case of error return is empty string with a customized error
func StripPort(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	strTemp := rxPort.FindString(u)
	if rxPort.Match([]byte(u)) {
		return strings.Replace(u, strTemp, "", -1), nil
	}
	return "", errors.New("Port Does not exist")
}

// TrimURL checks if the passed string URL have any spaces to left or right side
// Accepts URL as a string argument
// returns modified string
func TrimURL(u string) string {
	u = strings.TrimSpace(u)
	return u
}
