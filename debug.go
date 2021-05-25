// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package debug is a Sidero-specific library for including debugging facilities for developers in our products
// when they are compiled with sidero.debug build tag. They are not included by default.
// Also provides utils for detecting if the code was compiled with race build tag.
package debug

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"sort"
)

// LogFunc is a simplest logging function.
type LogFunc func(msg string)

// ListenAndServe runs debug server on given address.
//
// If debugging is enabled (debug.Enabled is true), the function blocks until fatal error in encountered
// (it is returned) or ctx is canceled (in that case, nil is returned).
//
// If debugging is disabled (debug.Enabled is false), the function immediately returns nil.
//
// Both cases can be handled with the following code:
//
//   if err := debug.ListenAndServe(ctx, addr, log); err != nil {
//       log.Fatal(err)
//   }
func ListenAndServe(ctx context.Context, addr string, log LogFunc) error {
	if !Enabled {
		return nil
	}

	log("starting debug server")

	for _, h := range handlers() {
		log(fmt.Sprintf("http://%s%s", addr, h))
	}

	s := &http.Server{
		Addr: addr,
	}

	// no real reason to do a graceful shutdown with s.Shutdown
	go func() {
		<-ctx.Done()

		_ = s.Close() //nolint:errcheck
	}()

	err := s.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

// handlers returns all patterns registered on the default HTTP server mux.
// Useful to check for unexpected handlers.
func handlers() []string {
	keys := reflect.ValueOf(http.DefaultServeMux).Elem().FieldByName("m").MapKeys()

	res := make([]string, len(keys))
	for i, v := range keys {
		res[i] = v.String()
	}

	sort.Strings(res)

	return res
}

func init() {
	if !Enabled {
		// explicitly disable memory profiling to save around 1.4MiB of memory
		runtime.MemProfileRate = 0

		return
	}

	// that's defaults, just make them explicit
	runtime.SetCPUProfileRate(100)
	runtime.MemProfileRate = 512 * 1024

	// https://github.com/DataDog/go-profiler-notes/blob/main/block.md#overhead
	runtime.SetBlockProfileRate(10000)

	// no science behind that value
	runtime.SetMutexProfileFraction(10000)
}
