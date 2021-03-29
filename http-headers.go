package main

import (
	"net/http"
)

func Canonical(s string) string {
	return http.CanonicalHeaderKey(s)
}
