// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
//go:build !sidero.debug

package debug //nolint:testpackage // to test unexported method

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebugOff(t *testing.T) {
	assert.Empty(t, handlers())

	switch pprofEnabled() {
	case true:
		assert.Equal(t, 524288, runtime.MemProfileRate)
	case false:
		assert.Equal(t, 0, runtime.MemProfileRate)
	}
}
