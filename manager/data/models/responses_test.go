package models

import (
	"strings"
	"testing"
)

func TestToJSON(t *testing.T) {
	flag := Flag{
		ID:          1,
		Key:         "yessir",
		DisplayName: "YesSir!",
		Sdkkey:      "dafdacd",
		Status:      true,
	}

	flags := []Flag{flag, flag}

	tests := []struct {
		name  string
		input interface{}
		err   error
		want  bool
	}{
		{
			name:  "one flag",
			input: flag,
			err:   nil,
			want:  true,
		},
		{
			name:  "multiple flags",
			input: flags,
			err:   nil,
			want:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			raw, err := ToJSON(flags)
			have := string(*raw)
			if err != tc.err {
				t.Fatalf("have error: %v\nwant error: %v", err, tc.err)
			} else if strings.Contains(have, "yessir") != tc.want {
				t.Fatalf("have: %v", have)
			}
		})
	}
}
