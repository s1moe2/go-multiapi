package middleware

import (
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

// Chain takes an array of middleware functions and a final handler
// and chains them in the same order as in the array.
// middlewareChain([]middleware{m1, m2, m3}, h) ==> m1(m2(m3(h)))
func Chain(funcs []Middleware, final http.Handler) http.Handler {
	for i := range funcs {
		final = funcs[len(funcs)-1-i](final)
	}
	return final
}
