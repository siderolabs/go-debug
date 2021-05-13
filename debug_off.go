// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// +build !sidero.debug

package debug

// Enabled is false when compiled without sidero.debug build tag.
//
// Profiling is disabled to safe resources as a side-effect of importing this package.
const Enabled = false
