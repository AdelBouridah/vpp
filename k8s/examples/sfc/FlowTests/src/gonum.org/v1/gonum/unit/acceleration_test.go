// Code generated by "go generate gonum.org/v1/gonum/unit; DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"fmt"
	"testing"
)

func TestAcceleration(t *testing.T) {
	for _, value := range []float64{-1, 0, 1} {
		var got Acceleration
		err := got.From(Acceleration(value).Unit())
		if err != nil {
			t.Errorf("unexpected error for %T conversion: %v", got, err)
		}
		if got != Acceleration(value) {
			t.Errorf("unexpected result from round trip of %T(%v): got: %v want: %v", got, float64(value), got, value)
		}
		if got != got.Acceleration() {
			t.Errorf("unexpected result from self interface method call: got: %#v want: %#v", got, value)
		}
		err = got.From(ether(1))
		if err == nil {
			t.Errorf("expected error for ether to %T conversion", got)
		}
	}
}

func TestAccelerationFormat(t *testing.T) {
	for _, test := range []struct {
		value  Acceleration
		format string
		want   string
	}{
		{1.23456789, "%v", "1.23456789 m s^-2"},
		{1.23456789, "%.1v", "1 m s^-2"},
		{1.23456789, "%20.1v", "            1 m s^-2"},
		{1.23456789, "%20v", "   1.23456789 m s^-2"},
		{1.23456789, "%1v", "1.23456789 m s^-2"},
		{1.23456789, "%#v", "unit.Acceleration(1.23456789)"},
		{1.23456789, "%s", "%!s(unit.Acceleration=1.23456789 m s^-2)"},
	} {
		got := fmt.Sprintf(test.format, test.value)
		if got != test.want {
			t.Errorf("Format %q %v: got: %q want: %q", test.format, float64(test.value), got, test.want)
		}
	}
}