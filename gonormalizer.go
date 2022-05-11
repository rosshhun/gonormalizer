package gonormalizer

/*
foo://example.com:8042/over/there?name=ferret#nose
\_/   \______________/\_________/ \_________/ \__/
 |           |            |            |        |
scheme     authority       path        query   fragment
 |   _____________________|__
/ \ /                        \
urn:example:animal:ferret:nose
*/

/*
The following four URIs are equivalent:
  http://example.com
  http://example.com/
  http://example.com:/
  http://example.com:80/
*/
import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

const (
	// URL          = `^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`
	// HTTP_REGEXP = `^(?:https?:)?\/\/`
	URLPort     = `(:(\d{1,5}))`
	IP          = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	URL         = `^(([^:/?#]+):)?(//([^/?#]*))?([^?#]*)(\?([^#]*))?(#(.*))?`
	HTTP_REGEXP = `^(?:[a-zA-Z.-]+:)?\/\/`
	WWW_REGEXP  = `(www\.)`
	// GPROTO_REGEXP  = `(?:([a-zA-Z.-]+):)?\/\/`
	FHTTP_REGEXP   = `^(?:https:)`
	FHTTPS_REGEXP  = `^(?:http:)`
	SHASH_REGEXP   = `(#.+)`
	SAUTH_REGEXP   = `^((?:\w+:)?\/\/)?[^@/]+@`
	TAIL_REGEXP    = `\/$`
	TAILDOT_REGEXP = `\.$`
	COLON_REGEXP   = `:$`
)

// URL complete regular expression
var (
	rxC *regexp.Regexp = regexp.MustCompile(URL)

	// https or http
	rxHttp *regexp.Regexp = regexp.MustCompile(HTTP_REGEXP)

	// tail slash
	rxSlash *regexp.Regexp = regexp.MustCompile(TAIL_REGEXP)

	//tail dot
	rxTd *regexp.Regexp = regexp.MustCompile(TAILDOT_REGEXP)

	// www
	rxWWW *regexp.Regexp = regexp.MustCompile(WWW_REGEXP)

	// force http
	rxFhttp *regexp.Regexp = regexp.MustCompile(FHTTP_REGEXP)

	// force https
	rxFhttps *regexp.Regexp = regexp.MustCompile(FHTTPS_REGEXP)

	// strip hash
	rxShash *regexp.Regexp = regexp.MustCompile(SHASH_REGEXP)

	// strip auth
	rxSauth *regexp.Regexp = regexp.MustCompile(SAUTH_REGEXP)

	// golbal protocols
	// rxGProto *regexp.Regexp = regexp.MustCompile(GPROTO_REGEXP)

	rxPort *regexp.Regexp = regexp.MustCompile(URLPort)

	rxColon *regexp.Regexp = regexp.MustCompile(COLON_REGEXP)
)

func AddPort(u string, p string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	if rxSlash.Match([]byte(u)) {
		u, _ = RemoveTrailingSlash(u)
	}
	if !rxPort.Match([]byte(u)) {
		return u + ":" + p, nil
	}
	return "", errors.New("Port already exist")
}

// add a default protocol
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

// to add ending slash
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

// to add tail dot
func AddTrailingDot(u string) (string, error) {
	if !IsValid(u) && !IsEmpty(u) {
		return "", errors.New("Not a vaild URL")
	}
	u = TrimURL(u)
	u = LowerCase(u)
	if rxSlash.Match([]byte(u)) {
		u, _ = RemoveTrailingSlash(u)
	}
	return u + ".", nil
}

// add a default protocol
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

// convert url http
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

// convert url to https
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

// to check URL is  Valid
func IsValid(u string) bool {
	return rxC.Match([]byte(u))
}

func IsEmpty(u string) bool {
	u = strings.TrimSpace(u)
	return len(u) == 0
}

func LowerCase(u string) string {
	return strings.ToLower(u)
}

// to remove ending slash
func RemoveTrailingSlash(u string) (string, error) {
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

// to remove tail dot
func RemoveTrailingDot(u string) (string, error) {
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

// Protocol used for getting the Protocol of the URL
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

// to remove http or https or //
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

// to remove www from URL, NOT WORKING
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

// strip # from url
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

// strip #:~:text=hello from url
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

//strip authentication part of a url
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

func TrimURL(u string) string {
	u = strings.TrimSpace(u)
	return u
}
