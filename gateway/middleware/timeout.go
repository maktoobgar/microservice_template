package middleware

import (
	"context"
	"net/http"
	"time"

	"service/pkg/errors"
)

func Timeout(timeout time.Duration, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new context with the given timeout
		ctx, cancel := context.WithTimeout(r.Context(), timeout*time.Second)
		defer cancel() // Cancel the context to release resources when the request completes

		done := make(chan struct{})
		go func() {
			Panic(next).ServeHTTP(w, r.WithContext(ctx))
			close(done)
		}()
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				panic(errors.New(errors.TimeoutStatus, errors.Resend, "TimeoutError"))
			}
		case <-done:
			return
		}
	})
}
