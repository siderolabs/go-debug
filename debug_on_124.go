// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
//go:build go1.24 && !go1.26

package debug

import (
	"net/http"
	"reflect"
	"slices"
)

// handlers returns all patterns registered on the default HTTP server mux.
// Useful to check for unexpected handlers.
func handlers() []string {
	res := make([]string, 0, 10)
	index := reflect.ValueOf(http.DefaultServeMux).Elem().FieldByName("index")
	segmentsMap := index.FieldByName("segments")
	segmentsRange := segmentsMap.MapRange()

	for segmentsRange.Next() {
		slc := segmentsRange.Value()

		for i := range slc.Len() {
			res = append(res, slc.Index(i).Elem().FieldByName("str").String())
		}
	}

	multis := index.FieldByName("multis")
	for i := range multis.Len() {
		res = append(res, multis.Index(i).Elem().FieldByName("str").String())
	}

	slices.Sort(res)

	return slices.Compact(res)
}
