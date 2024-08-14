// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
//go:build go1.22 && !go1.24

package debug

import (
	"net/http"
	"reflect"
	"slices"
)

// handlers returns all patterns registered on the default HTTP server mux.
// Useful to check for unexpected handlers.
func handlers() []string {
	// This will only work for Go 1.22, since `http.ServeMux.patterns` contains "remove if possible" comment string
	// from Go developers.
	patterns := reflect.ValueOf(http.DefaultServeMux).Elem().FieldByName("patterns")
	pLen := patterns.Len()

	res := make([]string, pLen)
	for i := range pLen {
		res[i] = patterns.Index(i).Elem().FieldByName("str").String()
	}

	slices.Sort(res)

	return res
}
