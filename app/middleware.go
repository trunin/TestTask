package app

import (
	"fmt"
	"net/http"
)

const maxHttpRequests = 100

func (s *Server) requestLimits(h http.Handler) http.Handler {
	sema := make(chan struct{}, maxHttpRequests)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case sema <- struct{}{}:
			h.ServeHTTP(w, r)
			<-sema
		default:
			fmt.Fprint(w, "max requests processing")
		}
	})
}