package version

import "testing"

func TestChange(t *testing.T) {
  testCases := []struct {
    in    string
    vtype string
    out   string
  }{
    {"0.4", "major", "1.0.0"},
    {"0.4.0-hash", "major", "1.0.0"},
    {"1.0", "major", "2.0.0"},
    {"1", "major", "2.0.0"},
    {"1.0.1", "minor", "1.1.0"},
  }

  var err error
  v := Version{}
  for _, tt := range testCases {
    err = v.Change(tt.in, tt.vtype)
    if err != nil {
      t.Fatal(err)
    }
    if v.String() != tt.out {
      t.Errorf("Change(%s,%s): got %s, want %s", tt.in, tt.vtype, v.String(), tt.out)
    }
  }
}

func TestInc(t *testing.T) {
  testCases := []struct {
    vtype string
    out   string
  }{
    {"major", "1.0.0"},
    {"major", "2.0.0"},
    {"patch", "2.0.1"},
    {"minor", "2.1.0"},
    {"patch", "2.1.1"},
    {"major", "3.0.0"},
  }

  var err error
  v := Version{}
  v.Set("0.1.2-hash")
  for _, tt := range testCases {
    err = v.Inc(tt.vtype)
    if err != nil {
      t.Fatal(err)
    }
    if v.String() != tt.out {
      t.Errorf("Inc(%s): got %s, want %s", tt.vtype, v.String(), tt.out)
    }
  }
}

func TestIsValid(t *testing.T) {
  testCases := []struct {
    in string
    out bool
  }{
    {"1", true},
    {"5.2.22-hash", true},
    {"-hash", false},
    {"", false},
    {"2.12-hash", true},
    {"1.2.3", true},
    {"1-hash", true},
  }

  var isvalid bool
  for _, tt := range testCases {
    isvalid, _ = IsValid(tt.in)
    if isvalid != tt.out {
      t.Errorf("IsValid(%s): got %t, want %t", tt.in, isvalid, tt.out)
    }
  }
}
