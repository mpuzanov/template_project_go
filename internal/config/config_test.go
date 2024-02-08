package config

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {

	type testCase struct {
		test        string
		want        string
		expectedErr error
	}
	// Create new test cases
	testCases := []testCase{
		{
			test:        "is ok version",
			want:        "development",
			expectedErr: nil,
		}, {
			test:        "is not ok version",
			want:        "not development",
			expectedErr: errors.New("not development"),
		},
	}

	for _, tc := range testCases {
		// Run Tests
		t.Run(tc.test, func(t *testing.T) {
			got := ShowVersion()
			if tc.expectedErr != nil {
				require.NotEqual(t, tc.want, got, tc.test)
			} else {
				require.Equal(t, tc.want, got, tc.test)
				//t.Errorf("got %q want %q", got, tc.want)
			}
		})
	}
}

func TestVersion2(t *testing.T) {
	got := ShowVersion()
	want := "development"

	require.Equal(t, want, got)
	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }
}
