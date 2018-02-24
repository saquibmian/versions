package versions

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	versionSeparatorCharacter = "-"
)

var (
	// ErrInvalidVersionFormat is returned from ParseString when the version string is of an invalid format.
	ErrInvalidVersionFormat = fmt.Errorf("versions: invalid version format")
)

var comparisonReturnValue = map[bool]int{
	true:  1,
	false: -1,
}

type Version struct {
	Major  int
	Minor  int
	Patch  int
	Build  int
	Suffix string
}

func (v Version) String() string {
	var suffix string
	if v.Suffix != "" {
		suffix = versionSeparatorCharacter + v.Suffix
	}

	return fmt.Sprintf(
		"%d.%d.%d.%d%s",
		v.Major,
		v.Minor,
		v.Patch,
		v.Build,
		suffix,
	)
}

// Compare compares two versions.
// Major, Minor, Patch, and Build are compared in order; finally, Suffix is lexicographically compared.
func Compare(v1 Version, v2 Version) int {
	if v1.Major != v2.Major {
		return comparisonReturnValue[v1.Major > v2.Major]
	}
	if v1.Minor != v2.Minor {
		return comparisonReturnValue[v1.Minor > v2.Minor]
	}
	if v1.Patch != v2.Patch {
		return comparisonReturnValue[v1.Patch > v2.Patch]
	}
	if v1.Build != v2.Build {
		return comparisonReturnValue[v1.Build > v2.Build]
	}

	return strings.Compare(v1.Suffix, v2.Suffix)
}

func Equal(v1 Version, v2 Version) bool {
	return Compare(v1, v2) == 0
}

// ParseString parses a string-encoded Version.
// The following encodings are supported:
// - 1
// - 1-suffix
// - 1.0
// - 1.0-suffix
// - 1.0.0
// - 1.0.0-suffix
// - 1.0.0.0
// - 1.0.0.0-suffix
// All version numbers are 32-bit integers; numbers larger that, or incorrectly formatted version strings will
// return ErrInvalidVersionFormat.
func ParseString(str string) (Version, error) {
	var toReturn Version

	versionParts := strings.Split(str, versionSeparatorCharacter)
	if len(versionParts) < 1 || len(versionParts) > 2 {
		return toReturn, ErrInvalidVersionFormat
	}

	digitParts := strings.Split(versionParts[0], ".")
	for i, numberString := range digitParts {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			return toReturn, ErrInvalidVersionFormat
		}
		switch i {
		case 0:
			toReturn.Major = number
		case 1:
			toReturn.Minor = number
		case 2:
			toReturn.Patch = number
		case 3:
			toReturn.Build = number
		default:
			return toReturn, ErrInvalidVersionFormat
		}
	}

	if len(versionParts) == 2 {
		toReturn.Suffix = versionParts[1]
	}

	return toReturn, nil
}

type AscendingVersions []Version

type DescendingVersions []Version

// NewVersion creates a new version with the supplied numbers and suffix.
func NewVersion(major int, minor int, patch int, build int, suffix string) Version {
	return Version{major, minor, patch, build, suffix}
}
