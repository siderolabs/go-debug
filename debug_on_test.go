// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
//go:build sidero.debug

package debug //nolint:testpackage // to test unexported method

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebugOn(t *testing.T) {
	expected := []string{
		"/debug/events",
		"/debug/pprof/",
		"/debug/pprof/cmdline",
		"/debug/pprof/profile",
		"/debug/pprof/symbol",
		"/debug/pprof/trace",
		"/debug/requests",
		"/debug/vars",
	}
	assert.Equal(t, expected, handlers())

	assert.Equal(t, 524288, runtime.MemProfileRate)
}
