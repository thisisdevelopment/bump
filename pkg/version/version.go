package version

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Type defines the semantic version type
type Type string

const (
	// Major indicates breaking changes
	Major = Type("major")
	// Minor indicates non-breaking changes
	Minor = Type("minor")
	// Patch indicates no breaking changes
	Patch = Type("patch")
)

// Version as in semantic versioning
type Version struct {
	Major int64
	Minor int64
	Patch int64
}

func (v *Version) String() string {
	switch true {
	case v.Major >= 0 && v.Minor >= 0 && v.Patch >= 0:
		return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	case v.Major >= 0 && v.Minor >= 0:
		return fmt.Sprintf("%d.%d", v.Major, v.Minor)
	case v.Major >= 0:
		return fmt.Sprintf("%d", v.Major)
	default:
		return "%!s(INVALID_VERSION)"
	}
}

// Parse a version string of the forms "2", "2.3", or "0.10.11".
// Any information after the third number ("2.0.0-beta") is discarded.
// If a field is omitted from the string version (e.g. "0.2"), it's stored in
// the Version string as -1.
func Parse(version string) (*Version, error) {
	if len(version) == 0 {
		return nil, errors.New("version empty")
	}

	parts := strings.SplitN(version, ".", 3)
	if len(parts) == 1 {
		major, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return nil, err
		}
		return &Version{
			Major: major,
			Minor: -1,
			Patch: -1,
		}, nil
	}
	if len(parts) == 2 {
		major, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return nil, err
		}
		minor, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return nil, err
		}
		return &Version{
			Major: major,
			Minor: minor,
			Patch: -1,
		}, nil
	}
	if len(parts) == 3 {
		major, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return nil, err
		}
		minor, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return nil, err
		}
		patchParts := strings.SplitN(parts[2], "-", 2)
		patch, err := strconv.ParseInt(patchParts[0], 10, 64)
		if err != nil {
			return nil, err
		}
		return &Version{
			Major: major,
			Minor: minor,
			Patch: patch,
		}, nil
	}
	return nil, fmt.Errorf("invalid version: %s", version)
}

// Change takes a basic literal representing a string version, and
// increments the version number per the given VersionType.
func Change(vtype Type, value string) (*Version, error) {
	versionNoQuotes := strings.Replace(value, "\"", "", -1)
	version, err := Parse(versionNoQuotes)
	if err != nil {
		return nil, errors.Wrap(err, "parsing version")
	}
	switch true {
	case vtype == Major:
		version.Major++
		if version.Minor != -1 {
			version.Minor = 0
		}
		if version.Patch != -1 {
			version.Patch = 0
		}
	case vtype == Minor:
		if version.Minor == -1 {
			version.Minor = 0
		}
		if version.Patch != -1 {
			version.Patch = 0
		}
		version.Minor++
	case vtype == Patch:
		if version.Patch == -1 {
			version.Patch = 0
		}
		version.Patch++
	default:
		return nil, fmt.Errorf("invalid type: %s", vtype)
	}

	return version, nil
}
