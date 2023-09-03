package timetracedb

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Version struct {
	Meta  string `json:"meta"  xml:"meta"`
	Major uint8  `json:"major" xml:"major"`
	Minor uint8  `json:"minor" xml:"minor"`
	Patch uint8  `json:"patch" xml:"patch"`
}

var version = Version{
	Meta:  "beta",
	Major: 0,
	Minor: 1,
	Patch: 0,
}

func Agent() string {
	return fmt.Sprintf("timetrace/%s", GetVersionAsString())
}

func StringVersion() string {
	return fmt.Sprintf("%d.%d.%d-%s", version.Major, version.Minor, version.Patch, version.Meta)
}

func JSONVersion() (string, error) {
	v, err := json.Marshal(version)
	if err != nil {
		return "", err
	}

	return string(v), nil
}

func XMLVersion() (string, error) {
	v, err := xml.MarshalIndent(version, "", " ")
	if err != nil {
		return "", err
	}

	return string(v), nil
}
