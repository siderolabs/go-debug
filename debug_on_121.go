// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
//go:build !go1.22

package debug

import (
	"net/http"
	"reflect"
	"slices"
)

// handlers returns all patterns registered on the default HTTP server mux.
// Useful to check for unexpected handlers.
func handlers() []string {
	keys := reflect.ValueOf(http.DefaultServeMux).Elem().FieldByName("m").MapKeys()

	res := make([]string, len(keys))
	for i, v := range keys {
		res[i] = v.String()
	}

	slices.Sort(res)

	return res
}
