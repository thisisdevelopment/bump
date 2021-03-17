package version

import (
  "fmt"
  "github.com/pkg/errors"
)

const (
  major = iota
  minor
  patch
  format = "%d.%d.%d"
)

// Version as in semantic versioning
type Version struct {
  ver [3]int
}

var vtypes = map[string]int{"major": major, "minor": minor, "patch": patch}

// IsValid checks to see whether the version string would be accepted
// by the Set and Change methods, indicated by the bool return value. Note
// that the error value could be non-nil, and the boolean true. This can
// happen with a partial version string for example, which is allowed. So the
// error value is only relevant when the bool is false (= string is invalid)
func IsValid(ver string) (bool, error) {
  var i int
  i, err := fmt.Sscanf(ver, format, &i, &i, &i)
  return i > 0, err
}

// String makes sure a Version value is outputted as intended
func (v *Version) String() string {
  return fmt.Sprintf(format, v.ver[major], v.ver[minor], v.ver[patch])
}

// Set initializes/resets the Version value.
// Any hash value is stripped, i.e. "1.2.3-hash" becomes "1.2.3".
// Partial version strings are allowed, i.e. "1" becomes "1.0.0".
func (v *Version) Set(ver string) error {
  i, err := fmt.Sscanf(ver, format, &v.ver[major], &v.ver[minor], &v.ver[patch])

  // check if there was at least 1 integer found
  if i > 0 {
    // zero out lesser fields, if any
    for ;i < len(v.ver);i++ {
      v.ver[i] = 0
    }

    // partial version strings are allowed, so ignore err
    return nil
  }

  return errors.Wrap(err, "parsing version")
}

// Inc increments the indicated portion of the Version value
func (v *Version) Inc(vtype string) error {
  i, ok := vtypes[vtype]
  if ! ok {
    return fmt.Errorf("invalid type: %s", vtype)
  }

  v.ver[i]++

  // zero out lesser fields, if any
  for i++;i < len(v.ver);i++ {
    v.ver[i] = 0
  }

  return nil
}

// Change initializes/resets the Version value, and increments the indicated portion of the Version value in one go
func (v *Version) Change(ver, vtype string) error {
  err := v.Set(ver)
  if err != nil {
    return err
  }
  err = v.Inc(vtype)
  if err != nil {
    return err
  }
  return nil
}
