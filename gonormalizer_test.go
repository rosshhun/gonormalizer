package gonormalizer

import "testing"

func TestAddPort(t *testing.T) {
	var tests = []struct {
		param1   string
		param2   string
		expected string
	}{
		{"https://www.example.com", "80", "https://www.example.com:80"},
		{"https://www.example.com/", "80", "https://www.example.com:80"},
		{"https://example.com/", "80", "https://example.com:80"},
		{"https://example.com", "80", "https://example.com:80"},
		{"example.com", "80", "example.com:80"},
		{"example.com/", "80", "example.com:80"},
		{"example.com:/", "80", "example.com:80"},
	}
	for _, test := range tests {
		actual, _ := AddPort(test.param1, test.param2)
		if actual != test.expected {
			t.Errorf("Expected AddPort(%q,%q) to be %v, got %v", test.param1, test.param2, test.expected, actual)
		}
	}
}

func TestAddProtocol(t *testing.T) {
	var tests = []struct {
		param1   string
		param2   string
		expected string
	}{
		{"www.example.com", "https", "https://www.example.com"},
		{"example.com/", "http", "http://example.com/"},
		{"example.com/", "ftp", "ftp://example.com/"},
		{"example.com", "udp", "udp://example.com"},
		{"example.co.in", "https", "https://example.co.in"},
	}
	for _, test := range tests {
		actual, _ := AddProtocol(test.param1, test.param2)
		if actual != test.expected {
			t.Errorf("Expected AddProtocol(%q,%q) to be %v, got %v", test.param1, test.param2, test.expected, actual)
		}
	}
}

func TestAddTrailingSlash(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"https://www.example.com.in", "https://www.example.com.in/"},
		{"www.example.com", "www.example.com/"},
		{"example.com", "example.com/"},
		{"http://example.com", "http://example.com/"},
		{"user:password@example.com", "user:password@example.com/"},
	}
	for _, test := range tests {
		actual, _ := AddTrailingSlash(test.param)
		if actual != test.expected {
			t.Errorf("Expected AddTrailingSlash(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestAddTrailingDot(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"https://www.example.com.in", "https://www.example.com.in."},
		{"www.example.com/", "www.example.com."},
		{"example.com/", "example.com."},
		{"http://example.com", "http://example.com."},
		{"user:password@example.com", "user:password@example.com."},
	}
	for _, test := range tests {
		actual, _ := AddTrailingDot(test.param)
		if actual != test.expected {
			t.Errorf("Expected AddTrailingDot(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestDefaultProtocol(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"www.example.com.in", "http://www.example.com.in"},
		{"www.example.com/", "http://www.example.com/"},
		{"example.com/", "http://example.com/"},
		{"example.com.", "http://example.com."},
		{"user:password@example.com", "http://user:password@example.com"},
	}
	for _, test := range tests {
		actual, _ := DefaultProtocol(test.param)
		if actual != test.expected {
			t.Errorf("Expected DefaultProtocol(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestForceHttp(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"https://www.example.com.in", "http://www.example.com.in"},
		{"https://example.com", "http://example.com"},
		{"https://user:password@example.com", "http://user:password@example.com"},
	}
	for _, test := range tests {
		actual, _ := ForceHttp(test.param)
		if actual != test.expected {
			t.Errorf("Expected ForceHttp(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestForceHttps(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"http://www.example.com.in", "https://www.example.com.in"},
		{"http://example.com", "https://example.com"},
		{"http://user:password@example.com", "https://user:password@example.com"},
	}
	for _, test := range tests {
		actual, _ := ForceHttps(test.param)
		if actual != test.expected {
			t.Errorf("Expected ForceHttps(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsValid(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"http://www.example.com.in", true},
		{"http://example.com", true},
		{"https://example.com", true},
		{"https://example.com/", true},
		{"https://example.com:80/", true},
		{"ftp://example.com:80/", true},
		{"udp://example.com", true},
		{"", false},
		{"http://user:password@example.com", true},
	}
	for _, test := range tests {
		actual := IsValid(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsValid(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"http://www.example.com.in", false},
		{"", true},
	}
	for _, test := range tests {
		actual := IsEmpty(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsEmpty(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestLowerCase(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"http://WWW.example.com.in", "http://www.example.com.in"},
		{"https://Example.Com", "https://example.com"},
		{"httpS://user:password@example.com", "https://user:password@example.com"},
	}
	for _, test := range tests {
		actual := LowerCase(test.param)
		if actual != test.expected {
			t.Errorf("Expected LowerCase(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestStripTrailingSlash(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"http://www.example.com.in/", "http://www.example.com.in"},
		{"example.com./", "example.com."},
		{"http://user:password@example.com/", "http://user:password@example.com"},
	}
	for _, test := range tests {
		actual, _ := StripTrailingSlash(test.param)
		if actual != test.expected {
			t.Errorf("Expected StripTrailingSlash(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestStripTrailingDot(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"http://www.example.com.in.", "http://www.example.com.in"},
		{"example.com.", "example.com"},
		{"http://user:password@example.com/.", "http://user:password@example.com/"},
	}
	for _, test := range tests {
		actual, _ := StripTrailingDot(test.param)
		if actual != test.expected {
			t.Errorf("Expected StripTrailingDot(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestScheme(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"http://www.example.com.in.", "http"},
		{"udp://example.com.", "udp"},
		{"ftp://user:password@example.com/.", "ftp"},
	}
	for _, test := range tests {
		actual, _ := Scheme(test.param)
		if actual != test.expected {
			t.Errorf("Expected Scheme(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestStripProtocol(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"http://www.example.com.in.", "www.example.com.in."},
		{"udp://example.com", "example.com"},
		{"ftp://user:password@example.com", "user:password@example.com"},
	}
	for _, test := range tests {
		actual, _ := StripProtocol(test.param)
		if actual != test.expected {
			t.Errorf("Expected StripProtocol(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestStripWWW(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"http://www.example.com.in.", "http://example.com.in."},
		{"udp://www.example.com", "udp://example.com"},
	}
	for _, test := range tests {
		actual, _ := StripWWW(test.param)
		if actual != test.expected {
			t.Errorf("Expected StripWWW(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestStripHash(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"http://www.example.com.in#content", "http://www.example.com.in"},
		{"http://example.com#", "http://example.com"},
	}
	for _, test := range tests {
		actual, _ := StripHash(test.param)
		if actual != test.expected {
			t.Errorf("Expected StripHash(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestStripTextFragment(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"http://www.example.com.in#content", "http://www.example.com.in"},
		{"http://example.com#!#@~", "http://example.com"},
	}
	for _, test := range tests {
		actual, _ := StripTextFragment(test.param)
		if actual != test.expected {
			t.Errorf("Expected StripTextFragment(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestStripAuthentication(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"ftp://user:password@example.com", "ftp://example.com"},
		{"user:password@example.com", "example.com"},
	}
	for _, test := range tests {
		actual, _ := StripAuthentication(test.param)
		if actual != test.expected {
			t.Errorf("Expected StripAuthentication(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestStripPort(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"https://www.example.com:80", "https://www.example.com"},
		{"https://example.com:80", "https://example.com"},
		{"https://example.com:80", "https://example.com"},
		{"example.com:80", "example.com"},
		{"example.com:80/", "example.com/"},
	}
	for _, test := range tests {
		actual, _ := StripPort(test.param)
		if actual != test.expected {
			t.Errorf("Expected StripPort(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestTrimURL(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"https://www.example.com   ", "https://www.example.com"},
		{"         https://example.com", "https://example.com"},
		{"      https://example.com      ", "https://example.com"},
		{"    example.com ", "example.com"},
		{"   example.com/     ", "example.com/"},
	}
	for _, test := range tests {
		actual := TrimURL(test.param)
		if actual != test.expected {
			t.Errorf("Expected TrimURL(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}
