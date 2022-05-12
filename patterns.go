package gonormalizer

import "regexp"

// Basic regular expressions for validating strings
const (
	URLPort        = `(:(\d{1,5}))`
	URL            = `^(([^:/?#]+):)?(//([^/?#]*))?([^?#]*)(\?([^#]*))?(#(.*))?`
	HTTP_REGEXP    = `^(?:[a-zA-Z.-]+:)?\/\/`
	WWW_REGEXP     = `(www\.)`
	FHTTP_REGEXP   = `^(?:https:)`
	FHTTPS_REGEXP  = `^(?:http:)`
	SHASH_REGEXP   = `(#.+)`
	SAUTH_REGEXP   = `^((?:\w+:)?\/\/)?[^@/]+@`
	TAIL_REGEXP    = `\/$`
	TAILDOT_REGEXP = `\.$`
	COLON_REGEXP   = `:$`
)

// Variables storing compiled version of regular expressions
var (
	rxC      *regexp.Regexp = regexp.MustCompile(URL)
	rxHttp   *regexp.Regexp = regexp.MustCompile(HTTP_REGEXP)
	rxSlash  *regexp.Regexp = regexp.MustCompile(TAIL_REGEXP)
	rxTd     *regexp.Regexp = regexp.MustCompile(TAILDOT_REGEXP)
	rxWWW    *regexp.Regexp = regexp.MustCompile(WWW_REGEXP)
	rxFhttp  *regexp.Regexp = regexp.MustCompile(FHTTP_REGEXP)
	rxFhttps *regexp.Regexp = regexp.MustCompile(FHTTPS_REGEXP)
	rxShash  *regexp.Regexp = regexp.MustCompile(SHASH_REGEXP)
	rxSauth  *regexp.Regexp = regexp.MustCompile(SAUTH_REGEXP)
	rxPort   *regexp.Regexp = regexp.MustCompile(URLPort)
	rxColon  *regexp.Regexp = regexp.MustCompile(COLON_REGEXP)
)
