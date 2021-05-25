# go-debug

[![Go Reference](https://pkg.go.dev/badge/github.com/talos-systems/go-debug.svg)](https://pkg.go.dev/github.com/talos-systems/go-debug)

`go-debug` is a Sidero-specific library for including debugging facilities for developers in our products when they are compiled with `sidero.debug` build tag.
They are not included by default.
Also provides utils for detecting if the code was compiled with `race` build tag.
