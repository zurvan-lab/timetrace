package timetracedb

import "fmt"

// These constants follow the semantic versioning 2.0.0 spec (http://semver.org/)
var (
	major = 0
	minor = 1
	patch = 0
	meta  = "beta"
)

func StringVersion() string {
	v := fmt.Sprintf("%d.%d.%d", major, minor, patch)

	if meta != "" {
		v = fmt.Sprintf("%s-%s", v, meta)
	}

	return v
}
