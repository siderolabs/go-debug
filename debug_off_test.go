// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// +build !sidero.debug

package debug //nolint:testpackage // to test unexported method

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebugOff(t *testing.T) {
	expected := []string{}
	assert.Equal(t, expected, handlers())
}
