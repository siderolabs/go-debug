// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
//go:build sidero.debug && go1.23 && !go1.24

package debug //nolint:testpackage // to test unexported method

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebugOn(t *testing.T) {
	expected := []string{
		"/debug/events",
		"/debug/requests",
		"GET /debug/pprof/",
		"GET /debug/pprof/cmdline",
		"GET /debug/pprof/profile",
		"GET /debug/pprof/symbol",
		"GET /debug/pprof/trace",
		"GET /debug/vars",
	}
	assert.Equal(t, expected, handlers())

	assert.Equal(t, 524288, runtime.MemProfileRate)
}
