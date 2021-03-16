package version

import (
  "fmt"
  "github.com/pkg/errors"
)

const (
  major = iota
  minor
  patch
)

// Version as in semantic versioning
type Version struct {
  ver [3]int
}

var vtypes = map[string]int{"major": major, "minor": minor, "patch": patch}

func (v *Version) String() string {
  return fmt.Sprintf("%d.%d.%d", v.ver[major], v.ver[minor], v.ver[patch])
}

func (v *Version) Set(ver string) error {
  i, err := fmt.Sscanf(ver, "%d.%d.%d", &v.ver[major], &v.ver[minor], &v.ver[patch])
  if err != nil {
    return errors.Wrap(err, "parsing version")
  }

  for ;i < len(v.ver);i++ {
    v.ver[i] = 0
  }

  return nil
}

func (v *Version) Inc(vtype string) error {
  i, ok := vtypes[vtype]
  if ! ok {
    return fmt.Errorf("invalid type: %s", vtype)
  }

  v.ver[i]++

  for i++;i < len(v.ver);i++ {
    v.ver[i] = 0
  }

  return nil
}
