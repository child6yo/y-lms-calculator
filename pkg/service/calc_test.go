package service

import "testing"

func TestCulc(t *testing.T) {
	testCases := []struct {
		expression string
		expect     float64
		error      bool
	}{
		{
			"2+2*2",
			6,
			false,
		},
		{
			"(2+2)*2",
			8,
			false,
		},
		{
			"2*0",
			0,
			false,
		},
		{
			"2/0",
			0,
			true,
		},
		{
			"",
			0,
			true,
		},
	}

	for i, tc := range testCases {
		n, err := Calc(tc.expression)
		if err != nil {
			if tc.error {
				continue
			} else {
				t.Errorf("test %d failed: unexpected error (%s)", i+1, err)
			}
		}

		if n != tc.expect {
			t.Errorf("test %d failed: result %.2f, expect %.2f", i+1, n, tc.expect)
		}
	}
}