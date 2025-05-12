package handler

import (
	"log"
	"net/http"
	"net/http/httputil"

	"load-balancer/balancer"
)

// NewProxyHandler возвращает http.HandlerFunc
func NewProxyHandler(lb *balancer.LoadBalancer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		backend := lb.GetNextBackend()
		if backend == nil {
			http.Error(w, "No alive backends available", http.StatusServiceUnavailable)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(backend.URL)

		// Логируем
		log.Printf("Forwarding request to: %s", backend.URL.String())

		// Настраиваем обратный прокси
		r.Host = backend.URL.Host
		proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("Error proxying to backend %s: %v", backend.URL.String(), err)
			backend.SetAlive(false) // помечаем как мертвый
			http.Error(w, "Backend unavailable", http.StatusBadGateway)
		}

		proxy.ServeHTTP(w, r)
	}
}
