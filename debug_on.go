// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Protect with the build tag to avoid unwanted side-effects of imports below by default.
//go:build sidero.debug

package debug

import (
	_ "expvar"         // registers /debug/vars on http.DefaultServeMux
	_ "net/http/pprof" // registers /debug/pprof on http.DefaultServeMux

	// imported by google.golang.org/grpc anyway
	_ "golang.org/x/net/trace" // registers /debug/events and /debug/requests on http.DefaultServeMux
)

// Enabled is true when compiled with sidero.debug build tag.
//
// Profiling rates are configured as a side-effect of importing this package.
const Enabled = true
