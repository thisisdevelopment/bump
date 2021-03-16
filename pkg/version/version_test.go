package version

import "testing"

func TestChangeVersion(t *testing.T) {
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
